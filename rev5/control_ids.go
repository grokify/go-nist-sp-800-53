package rev5

type ControlIDs struct{}

func (ids ControlIDs) HighBaseline() []string     { return Rev5ProfileHighBaseline().ControlIDs() }
func (ids ControlIDs) ModerateBaseline() []string { return Rev5ProfileModerateBaseline().ControlIDs() }
func (ids ControlIDs) LowBaseline() []string      { return Rev5ProfileLowBaseline().ControlIDs() }

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
		"all":               len(Rev5Catalog().MustEnhancementIDs(IDTypeOSCAL, false)),
		"high_baseline":     len(ids.HighBaseline()),
		"high_uplift":       len(ids.HighUplift()),
		"moderate_baseline": len(ids.ModerateBaseline()),
		"moderate_uplift":   len(ids.ModerateUplift()),
		"low_baseline":      len(ids.LowBaseline())}
}
