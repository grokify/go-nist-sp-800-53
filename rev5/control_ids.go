package rev5

func ControlIDsHighBaseline() []string {
	c := Rev5ProfileHighBaseline()
	return c.ControlIDs()
}

func ControlIDsModerateBaseline() []string {
	c := Rev5ProfileModerateBaseline()
	return c.ControlIDs()
}

func ControlIDsLowBaseline() []string {
	c := Rev5ProfileLowBaseline()
	return c.ControlIDs()
}

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
