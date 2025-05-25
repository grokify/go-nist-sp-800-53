package rev5

import "fmt"

type ControlIDs struct{}

func (ids ControlIDs) HighBaseline() []string     { return ProfileHighBaseline().ControlIDs() }
func (ids ControlIDs) ModerateBaseline() []string { return ProfileModerateBaseline().ControlIDs() }
func (ids ControlIDs) LowBaseline() []string      { return ProfileLowBaseline().ControlIDs() }

func (ids ControlIDs) HighUplift() []string {
	return uplift(ids.HighBaseline(), ids.ModerateBaseline())
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
