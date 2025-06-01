package rev5

const (
	ClassSP80053Enhancement = "SP800-53-enhancement"

	GroupClassFamily = "family"

	ParamHowManyOneOrMore = "one-or-more"

	PropNameAltIdentifier = "alt-identifier"
	PropNameClass         = "class"
	PropNameLabel         = "label"
	PropClassSP80053      = "sp800-53a"

	// Risk Impact Tiers
	TierHighBaseline     = "high_baseline"
	TierHighUplift       = "high_uplift"
	TierModerateBaseline = "moderate_baseline"
	TierModerateUplift   = "moderate_uplift"
	TierLowBaseline      = "low_baseline"

	SepLines = "\n\n"
)

func Tiers() []string {
	return []string{
		TierHighBaseline,
		TierHighUplift,
		TierModerateBaseline,
		TierModerateUplift,
		TierLowBaseline}
}

func TierDisplayOrDefault(tierName, def string) string {
	m := map[string]string{
		TierHighBaseline:     "High Baseline",
		TierHighUplift:       "High Uplift",
		TierModerateBaseline: "Moderate Baseline",
		TierModerateUplift:   "Moderate Uplift",
		TierLowBaseline:      "Low Baseline",
	}
	if v, ok := m[tierName]; ok {
		return v
	} else {
		return def
	}
}
