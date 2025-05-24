package rev5

import (
	_ "embed"
	"encoding/json"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

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

func Rev5Catalog() *Catalog                 { return MustParseCatalogJSON(rev5Catalog) }
func Rev5CatalogHighBaseline() *Catalog     { return MustParseCatalogJSON(rev5CatalogHighBaseline) }
func Rev5CatalogModerateBaseline() *Catalog { return MustParseCatalogJSON(rev5CatalogModerateBaseline) }
func Rev5CatalogLowBaseline() *Catalog      { return MustParseCatalogJSON(rev5CatalogLowBaseline) }
func Rev5ProfileHighBaseline() *Profile     { return MustParseProfileJSON(rev5ProfileHighBaseline) }
func Rev5ProfileModerateBaseline() *Profile { return MustParseProfileJSON(rev5ProfileModerateBaseline) }
func Rev5ProfileLowBaseline() *Profile      { return MustParseProfileJSON(rev5ProfileLowBaseline) }

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
