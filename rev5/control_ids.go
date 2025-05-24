package rev5

func ControlIDsHighBaseline() []string     { return Rev5ProfileHighBaseline().ControlIDs() }
func ControlIDsModerateBaseline() []string { return Rev5ProfileModerateBaseline().ControlIDs() }
func ControlIDsLowBaseline() []string      { return Rev5ProfileLowBaseline().ControlIDs() }

func ControlIDsHighUplift() []string {
	return firstOnly(ControlIDsHighBaseline(), ControlIDsModerateBaseline())
}

func ControlIDsModerateUplift() []string {
	return firstOnly(ControlIDsModerateBaseline(), ControlIDsLowBaseline())
}

func firstOnly(s1, s2 []string) []string {
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

func ControlIDCounts() map[string]int {
	return map[string]int{
		"all":               len(Rev5Catalog().MustEnhancementIDs(IDTypeOSCAL, false)),
		"high_baseline":     len(ControlIDsHighBaseline()),
		"high_uplift":       len(ControlIDsHighUplift()),
		"moderate_baseline": len(ControlIDsModerateBaseline()),
		"moderate_uplift":   len(ControlIDsModerateUplift()),
		"low_baseline":      len(ControlIDsLowBaseline()),
	}
}
