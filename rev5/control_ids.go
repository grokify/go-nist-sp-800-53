package rev5

import (
	"fmt"
	"slices"
)

type ControlIDs struct{}

func NewControlIDs() *ControlIDs {
	return &ControlIDs{}
}

func (ids ControlIDs) HighBaseline() []string     { return ProfileHighBaseline().ControlIDs() }
func (ids ControlIDs) ModerateBaseline() []string { return ProfileModerateBaseline().ControlIDs() }
func (ids ControlIDs) LowBaseline() []string      { return ProfileLowBaseline().ControlIDs() }

func (ids ControlIDs) HighUplift() []string {
	return uplift(ids.HighBaseline(), ids.ModerateBaseline())
}

func (ids ControlIDs) HighUpliftIDs() (IDs, error) {
	return ParseIDs(ids.HighUplift())
}

func (ids ControlIDs) ModerateUplift() []string {
	return uplift(ids.ModerateBaseline(), ids.LowBaseline())
}

func uplift(s1, s2 []string) []string {
	var out []string
	prev := map[string]int{}
	for _, s := range s2 {
		prev[s] = 1
	}
	for _, s1 := range s1 {
		if _, ok := prev[s1]; !ok {
			out = append(out, s1)
		}
	}
	return out
}

func (ids ControlIDs) Counts() map[string]int {
	return map[string]int{
		"all":                len(CatalogAll().MustEnhancementIDs(IDTypeOSCAL, false)),
		TierHighBaseline:     len(ids.HighBaseline()),
		TierHighUplift:       len(ids.HighUplift()),
		TierModerateBaseline: len(ids.ModerateBaseline()),
		TierModerateUplift:   len(ids.ModerateUplift()),
		TierLowBaseline:      len(ids.LowBaseline())}
}

func (ids ControlIDs) ControlProfileStatus(idOSCAL string) ControlProfileStatus {
	out := ControlProfileStatus{}
	out.ID = idOSCAL
	if hb := ids.HighBaseline(); slices.Contains(hb, idOSCAL) {
		out.HighBaseline = true
	}
	if hu := ids.HighUplift(); slices.Contains(hu, idOSCAL) {
		out.HighUplift = true
	}
	if mb := ids.ModerateBaseline(); slices.Contains(mb, idOSCAL) {
		out.ModerateBaseline = true
	}
	if mu := ids.ModerateUplift(); slices.Contains(mu, idOSCAL) {
		out.ModerateUplift = true
	}
	if lb := ids.LowBaseline(); slices.Contains(lb, idOSCAL) {
		out.LowBaseline = true
	}
	return out
}

func (ids ControlIDs) Tier(tierName string) ([]string, error) {
	switch tierName {
	case TierHighBaseline:
		return ids.HighBaseline(), nil
	case TierHighUplift:
		return ids.HighUplift(), nil
	case TierModerateBaseline:
		return ids.ModerateBaseline(), nil
	case TierModerateUplift:
		return ids.ModerateUplift(), nil
	case TierLowBaseline:
		return ids.LowBaseline(), nil
	default:
		return []string{}, fmt.Errorf("tier not found (%s)", tierName)
	}
}

func (ids ControlIDs) TierIDs(tierName string) (IDs, error) {
	if rawIDs, err := ids.Tier(tierName); err != nil {
		return IDs{}, err
	} else {
		return ParseIDs(rawIDs)
	}
}

func (ids ControlIDs) TierMap(tierName string, idType IDType) (map[string]string, error) {
	idsTier, err := ids.Tier(tierName)
	if err != nil {
		return nil, err
	}
	if idType != IDTypeOSCAL {
		if idsWip, err := ParseIDs(idsTier); err != nil {
			return map[string]string{}, err
		} else if idsTry, err := idsWip.FormatIDType(idType); err != nil {
			return map[string]string{}, err
		} else {
			idsTier = idsTry
		}
	}
	return sliceToMap(idsTier), nil
}
