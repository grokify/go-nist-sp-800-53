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
	h := ControlIDsHighBaseline()
	m := ControlIDsModerateBaseline()
	prev := map[string]int{}
	var hu []string
	for _, id := range m {
		prev[id] = 1
	}
	for _, id := range h {
		if _, ok := prev[id]; !ok {
			hu = append(hu, id)
		}
	}
	return hu
}

func ControlIDsModerateUplift() []string {
	m := ControlIDsModerateBaseline()
	l := ControlIDsLowBaseline()
	prev := map[string]int{}
	var hu []string
	for _, id := range l {
		prev[id] = 1
	}
	for _, id := range m {
		if _, ok := prev[id]; !ok {
			hu = append(hu, id)
		}
	}
	return hu
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
