package rev5

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
	"github.com/opensecurity-dev/go-nist-sp-800-53/rev5/rev5tmpl"
)

type IDType string

const (
	IDTypeOSCAL     IDType = "id"
	IDTypeOSCALSort IDType = "sort-id"
	IDTypeNIST      IDType = "label"
	IDTypeNISTSort  IDType = "sp800-53a"

	PropNameStatus  = "status"
	StatusWithdrawn = "withdrawn"

	PartNameStatement = "statement"
	PartNameGuidance  = "guidance"
)

type Control oscalTypes.Control

func (ctr Control) ControlID() (ID, error) {
	return ParseID(ctr.ID)
}

func (ctr Control) Description() (string, error) {
	desc := ""
	if ctr.Parts != nil {
		parts := Parts(*ctr.Parts)
		if len(parts) > 0 {
			desc = strings.TrimSpace(parts.ExtractProseString())
		}
	}
	return ctr.renderParams(desc)
}

// Statement returns the control statement, with parameters rendered and
// items prefixed by their labels ("a.", "1."). It is empty for withdrawn
// controls.
func (ctr Control) Statement() (string, error) {
	return ctr.partText(PartNameStatement)
}

// Guidance returns the supplemental guidance, with parameters rendered.
func (ctr Control) Guidance() (string, error) {
	return ctr.partText(PartNameGuidance)
}

func (ctr Control) partText(name string) (string, error) {
	if ctr.Parts == nil {
		return "", nil
	}
	for _, p := range *ctr.Parts {
		if p.Name == name {
			var b strings.Builder
			writePartText(&b, p, 0)
			return ctr.renderParams(strings.TrimSpace(b.String()))
		}
	}
	return "", nil
}

// writePartText writes a part's prose and its sub-parts as an indented,
// labelled outline.
func writePartText(b *strings.Builder, p oscalTypes.Part, depth int) {
	label := ""
	if p.Props != nil {
		for _, prop := range *p.Props {
			if prop.Name == PropNameLabel {
				label = prop.Value
			}
		}
	}
	text := strings.TrimSpace(p.Prose)
	if label != "" || text != "" {
		b.WriteString(strings.Repeat("  ", depth))
		if label != "" {
			b.WriteString(label + " ")
		}
		b.WriteString(text + "\n")
	}
	if p.Parts != nil {
		next := depth
		if label != "" {
			next++
		}
		for _, sub := range *p.Parts {
			writePartText(b, sub, next)
		}
	}
}

// renderParams replaces parameter references in s with their values.
func (ctr Control) renderParams(s string) (string, error) {
	if s == "" || !strings.Contains(s, rev5tmpl.TemplateVarInsertPrefix) {
		return s, nil
	}
	paramSet, err := ctr.ParameterSet()
	if err != nil {
		return "", err
	}
	out, err := paramSet.RenderTemplate(s)
	if err != nil {
		return "", err
	}
	if strings.Contains(out, rev5tmpl.TemplateVarInsertPrefix) {
		return out, errors.New("control template not fully rendered")
	}
	return out, nil
}

// IDForType returns the control id for a requested status type.
func (ctr Control) IDForType(idtype IDType) (string, error) {
	if idtype == IDTypeOSCAL || strings.TrimSpace(string(idtype)) == "" {
		return ctr.ID, nil
	} else if ctr.Props == nil {
		return "", fmt.Errorf("no data found for found idtype (%s)", string(idtype))
	}
	for _, prop := range *ctr.Props {
		switch idtype {
		case IDTypeOSCALSort:
			if prop.Name == string(IDTypeOSCALSort) {
				if v := strings.TrimSpace(prop.Value); v != "" {
					return v, nil
				} else {
					return "", fmt.Errorf("no data found for found idtype (%s) for id (%s)", string(idtype), ctr.ID)
				}
			}
		case IDTypeNIST:
			if prop.Name == PropNameLabel && strings.TrimSpace(prop.Class) == "" {
				if v := strings.TrimSpace(prop.Value); v != "" {
					return v, nil
				} else {
					return "", fmt.Errorf("no data found for found idtype (%s) for id (%s)", string(idtype), ctr.ID)
				}
			}
		case IDTypeNISTSort:
			if prop.Name == PropNameLabel && prop.Class == string(IDTypeNISTSort) {
				if v := strings.TrimSpace(prop.Value); v != "" {
					return v, nil
				} else {
					return "", fmt.Errorf("no data found for found idtype (%s) for id (%s)", string(idtype), ctr.ID)
				}
			}
		}
	}
	return "", fmt.Errorf("no key found for found idtype (%s)", string(idtype))
}

func (ctr Control) HasImplementationLevelOrganization() bool {
	lvls := ctr.ImplementationLevels()
	return slices.Contains(lvls, PropImplementationLevelValueOrganization)
}

func (ctr Control) HasImplementationLevelSystem() bool {
	lvls := ctr.ImplementationLevels()
	return slices.Contains(lvls, PropImplementationLevelValueSystem)
}

func (ctr Control) ImplementationLevels() []string {
	var out []string
	if ctr.Props != nil {
		for _, prop := range *ctr.Props {
			if prop.Name == PropNameImplementationLevel {
				if v := strings.TrimSpace(prop.Value); v != "" {
					out = append(out, strings.ToLower(v)) // should all be lowercase already
				}
			}
		}
	}
	slices.Sort(out)
	return slices.Compact(out)
}

func (ctr Control) ParameterSet() (*ParameterSet, error) {
	return ParameterSetFromControl(&ctr)
}

func (ctr Control) Status() *string {
	if ctr.Props == nil {
		return nil
	}
	for _, prop := range *ctr.Props {
		if prop.Name == PropNameStatus {
			return &prop.Value
		}
	}
	return nil
}
