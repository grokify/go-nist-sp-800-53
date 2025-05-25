package rev5

import (
	_ "embed"
	"encoding/json"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

func CatalogAll() *Catalog                    { return MustParseCatalogJSON(rev5Catalog) }
func CatalogHighBaseline() *Catalog           { return MustParseCatalogJSON(rev5CatalogHighBaseline) }
func CatalogModerateBaseline() *Catalog       { return MustParseCatalogJSON(rev5CatalogModerateBaseline) }
func CatalogLowBaseline() *Catalog            { return MustParseCatalogJSON(rev5CatalogLowBaseline) }
func ControlSetHighBaseline() *ControlSet     { return mustControlSetTierName(TierHighBaseline) }
func ControlSetHighUplift() *ControlSet       { return mustControlSetTierName(TierHighUplift) }
func ControlSetModerateBaseline() *ControlSet { return mustControlSetTierName(TierModerateBaseline) }
func ControlSetModerateUplift() *ControlSet   { return mustControlSetTierName(TierModerateUplift) }
func ControlSetLowBaseline() *ControlSet      { return mustControlSetTierName(TierLowBaseline) }
func ProfileHighBaseline() *Profile           { return MustParseProfileJSON(rev5ProfileHighBaseline) }
func ProfileModerateBaseline() *Profile       { return MustParseProfileJSON(rev5ProfileModerateBaseline) }
func ProfileLowBaseline() *Profile            { return MustParseProfileJSON(rev5ProfileLowBaseline) }

//go:embed src/oscal_json_20240213_941c978/NIST_SP-800-53_rev5_catalog-min.json
var rev5Catalog []byte

//go:embed src/oscal_json_20240213_941c978/NIST_SP-800-53_rev5_HIGH-baseline-resolved-profile_catalog-min.json
var rev5CatalogHighBaseline []byte

//go:embed src/oscal_json_20240213_941c978/NIST_SP-800-53_rev5_MODERATE-baseline-resolved-profile_catalog-min.json
var rev5CatalogModerateBaseline []byte

//go:embed src/oscal_json_20240213_941c978/NIST_SP-800-53_rev5_LOW-baseline-resolved-profile_catalog-min.json
var rev5CatalogLowBaseline []byte

//go:embed src/oscal_json_20240213_941c978/NIST_SP-800-53_rev5_HIGH-baseline_profile-min.json
var rev5ProfileHighBaseline []byte

//go:embed src/oscal_json_20240213_941c978/NIST_SP-800-53_rev5_MODERATE-baseline_profile-min.json
var rev5ProfileModerateBaseline []byte

//go:embed src/oscal_json_20240213_941c978/NIST_SP-800-53_rev5_LOW-baseline_profile-min.json
var rev5ProfileLowBaseline []byte

// catalogWrapper embeds oscalTypes.Catalog for extension or additional methods.
type catalogWrapper struct {
	Catalog oscalTypes.Catalog `json:"catalog"`
}

func MustParseCatalogJSON(b []byte) *Catalog {
	w := catalogWrapper{}
	if err := json.Unmarshal(b, &w); err != nil {
		panic(err)
	}
	c := Catalog(w.Catalog)
	return &c
}

// profileWrapper embeds oscalTypes.Catalog for extension or additional methods.
type profileWrapper struct {
	Profile oscalTypes.Profile `json:"profile"`
}

func MustParseProfileJSON(b []byte) *Profile {
	w := profileWrapper{}
	if err := json.Unmarshal(b, &w); err != nil {
		panic(err)
	}
	p := Profile(w.Profile)
	return &p
}

func mustControlSetTierName(tierName string) *ControlSet {
	cat := CatalogAll()
	if ctrSet, err := cat.ControlSetTier(tierName); err != nil {
		panic(err)
	} else {
		return ctrSet
	}
}
