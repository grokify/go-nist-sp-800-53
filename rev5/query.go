package rev5

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

// Baseline names accepted by Query and reported in ControlInfo.Baselines.
const (
	BaselineLow      = "low"
	BaselineModerate = "moderate"
	BaselineHigh     = "high"
)

// Baselines returns the baseline names, lowest first.
func Baselines() []string {
	return []string{BaselineLow, BaselineModerate, BaselineHigh}
}

// ControlInfo is a flat summary of one control or enhancement, for listing
// and export.
type ControlInfo struct {
	// ID is the NIST form, such as "AC-2(1)".
	ID string `json:"id"`
	// OSCALID is the OSCAL form, such as "ac-2.1".
	OSCALID string `json:"oscalId"`
	// Family is the upper-case family abbreviation, such as "AC".
	Family      string `json:"family"`
	Title       string `json:"title"`
	Enhancement bool   `json:"enhancement"`
	Withdrawn   bool   `json:"withdrawn,omitempty"`
	// Baselines lists the baselines that include the control, lowest first.
	Baselines []string `json:"baselines"`
	// ImplementationLevels is "organization", "system" or both.
	ImplementationLevels []string `json:"implementationLevels,omitempty"`
	// Statement is the control statement with parameters rendered and
	// items labelled. It is set only by Catalog.ControlInfo, not by Query.
	Statement string `json:"statement,omitempty"`
	// Guidance is the supplemental guidance. It is set only by
	// Catalog.ControlInfo.
	Guidance string `json:"guidance,omitempty"`
}

// Query selects controls from a catalog. Empty fields match everything.
type Query struct {
	// Baseline is "low", "moderate" or "high".
	Baseline string
	// Families are family abbreviations, in any case.
	Families []string
	// Level is "organization" or "system".
	Level string
	// IncludeWithdrawn keeps withdrawn controls, which are left out by default.
	IncludeWithdrawn bool
}

// Validate reports an unknown baseline, family or level.
func (q Query) Validate() error {
	if q.Baseline != "" && !slices.Contains(Baselines(), q.Baseline) {
		return fmt.Errorf("unknown baseline %q: use %s", q.Baseline, strings.Join(Baselines(), ", "))
	}
	for _, f := range q.Families {
		if !IsFamilyAbbreviation(strings.ToLower(f)) {
			return fmt.Errorf("unknown family %q", f)
		}
	}
	switch q.Level {
	case "", PropImplementationLevelValueOrganization, PropImplementationLevelValueSystem:
	default:
		return fmt.Errorf("unknown implementation level %q: use %s or %s", q.Level,
			PropImplementationLevelValueOrganization, PropImplementationLevelValueSystem)
	}
	return nil
}

// baselineIndex maps baseline name to the set of OSCAL IDs it includes.
type baselineIndex map[string]map[string]bool

func newBaselineIndex() baselineIndex {
	ids := ControlIDs{}
	set := func(s []string) map[string]bool {
		m := make(map[string]bool, len(s))
		for _, id := range s {
			m[id] = true
		}
		return m
	}
	return baselineIndex{
		BaselineLow:      set(ids.LowBaseline()),
		BaselineModerate: set(ids.ModerateBaseline()),
		BaselineHigh:     set(ids.HighBaseline()),
	}
}

func (bi baselineIndex) of(oscalID string) []string {
	out := []string{}
	for _, b := range Baselines() {
		if bi[b][oscalID] {
			out = append(out, b)
		}
	}
	return out
}

// Query returns the controls matching q, in catalog order.
func (c *Catalog) Query(q Query) ([]ControlInfo, error) {
	if err := q.Validate(); err != nil {
		return nil, err
	}
	bi := newBaselineIndex()
	out := []ControlInfo{}
	for _, g := range c.Families(true) {
		if len(q.Families) > 0 && !slices.ContainsFunc(q.Families, func(f string) bool { return strings.EqualFold(f, g.ID) }) {
			continue
		}
		if g.Controls == nil {
			continue
		}
		for _, oc := range Controls(*g.Controls).Flatten() {
			info, err := summarize(Control(oc), bi)
			if err != nil {
				return nil, err
			}
			if info.Withdrawn && !q.IncludeWithdrawn {
				continue
			}
			if q.Baseline != "" && !slices.Contains(info.Baselines, q.Baseline) {
				continue
			}
			if q.Level != "" && !slices.Contains(info.ImplementationLevels, q.Level) {
				continue
			}
			out = append(out, info)
		}
	}
	return out, nil
}

// ControlInfo returns the summary of one control, including its statement
// and guidance. The ID may be in NIST or OSCAL form. If the statement or
// guidance cannot be fully rendered, it returns the summary with the error.
func (c *Catalog) ControlInfo(id string) (ControlInfo, error) {
	ctr, err := c.ControlByIDStringAny(id)
	if err != nil {
		return ControlInfo{}, err
	}
	info, err := summarize(*ctr, newBaselineIndex())
	if err != nil {
		return ControlInfo{}, err
	}
	var errs []error
	if info.Statement, err = ctr.Statement(); err != nil {
		errs = append(errs, fmt.Errorf("control %s statement: %w", info.ID, err))
	}
	if info.Guidance, err = ctr.Guidance(); err != nil {
		errs = append(errs, fmt.Errorf("control %s guidance: %w", info.ID, err))
	}
	return info, errors.Join(errs...)
}

func summarize(ctr Control, bi baselineIndex) (ControlInfo, error) {
	id, err := ctr.ControlID()
	if err != nil {
		return ControlInfo{}, err
	}
	nist, err := id.FormatNIST()
	if err != nil {
		return ControlInfo{}, err
	}
	s := ctr.Status()
	return ControlInfo{
		ID:                   nist,
		OSCALID:              ctr.ID,
		Family:               strings.ToUpper(id.FamilyAbbreviation),
		Title:                ctr.Title,
		Enhancement:          id.EnhancementNumber > 0,
		Withdrawn:            s != nil && *s == StatusWithdrawn,
		Baselines:            bi.of(ctr.ID),
		ImplementationLevels: ctr.ImplementationLevels(),
	}, nil
}

// IDCheck is the result of checking one control ID against a catalog.
type IDCheck struct {
	Input string `json:"input"`
	// ID is the NIST form, set when the input parses.
	ID string `json:"id,omitempty"`
	// Status is "ok", "withdrawn", "unknown" (well-formed but not in the
	// catalog) or "invalid" (not a control ID).
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

// ID check statuses.
const (
	IDCheckOK        = "ok"
	IDCheckWithdrawn = "withdrawn"
	IDCheckUnknown   = "unknown"
	IDCheckInvalid   = "invalid"
)

// CheckIDs reports whether each ID names a control in the catalog. IDs may be
// in NIST or OSCAL form.
func (c *Catalog) CheckIDs(inputs []string) []IDCheck {
	out := make([]IDCheck, 0, len(inputs))
	for _, in := range inputs {
		r := IDCheck{Input: in}
		id, err := ParseID(in)
		if err != nil {
			r.Status, r.Error = IDCheckInvalid, err.Error()
			out = append(out, r)
			continue
		}
		if r.ID, err = id.FormatNIST(); err != nil {
			r.Status, r.Error = IDCheckInvalid, err.Error()
			out = append(out, r)
			continue
		}
		ctr, err := c.Control(id)
		switch {
		case err != nil:
			r.Status, r.Error = IDCheckUnknown, err.Error()
		case ctr.Status() != nil && *ctr.Status() == StatusWithdrawn:
			r.Status = IDCheckWithdrawn
		default:
			r.Status = IDCheckOK
		}
		out = append(out, r)
	}
	return out
}

// FamilyInfo summarizes one control family.
type FamilyInfo struct {
	// ID is the upper-case abbreviation, such as "AC".
	ID    string `json:"id"`
	Title string `json:"title"`
	// Controls counts controls and enhancements that are not withdrawn.
	Controls int `json:"controls"`
	// Baselines counts the family's controls in each baseline.
	Baselines map[string]int `json:"baselines"`
}

// FamilyInfos summarizes every family, in catalog order.
func (c *Catalog) FamilyInfos() ([]FamilyInfo, error) {
	infos, err := c.Query(Query{})
	if err != nil {
		return nil, err
	}
	out := []FamilyInfo{}
	idx := map[string]int{}
	for _, g := range c.Families(false) {
		idx[strings.ToUpper(g.ID)] = len(out)
		out = append(out, FamilyInfo{ID: strings.ToUpper(g.ID), Title: g.Title, Baselines: map[string]int{}})
	}
	for _, info := range infos {
		i, ok := idx[info.Family]
		if !ok {
			continue
		}
		out[i].Controls++
		for _, b := range info.Baselines {
			out[i].Baselines[b]++
		}
	}
	return out, nil
}
