package rev5

const (
	ClassSP80053Enhancement = "SP800-53-enhancement"

	GroupClassFamily = "family"

	PropNameClass    = "class"
	PropNameLabel    = "label"
	PropClassSP80053 = "sp800-53a"

	// Risk Impact Tiers
	TierHighBaseline     = "high_baseline"
	TierHighUplift       = "high_uplift"
	TierModerateBaseline = "moderate_baseline"
	TierModerateUplift   = "moderate_uplift"
	TierLowBaseline      = "low_baseline"
)

func Tiers() []string {
	return []string{
		TierHighBaseline,
		TierHighUplift,
		TierModerateBaseline,
		TierModerateUplift,
		TierLowBaseline}
}
