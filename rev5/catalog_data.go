package rev5

import (
	_ "embed"
	"encoding/json"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

//go:embed src/oscal_json_20240213_941c978/NIST_SP-800-53_rev5_catalog-min.json
var rev5Catalog []byte

//go:embed src/oscal_json_20240213_941c978/NIST_SP-800-53_rev5_HIGH-baseline_profile-min.json
var rev5HighBaseline []byte

//go:embed src/oscal_json_20240213_941c978/NIST_SP-800-53_rev5_MODERATE-baseline_profile-min.json
var rev5ModerateBaseline []byte

//go:embed src/oscal_json_20240213_941c978/NIST_SP-800-53_rev5_LOW-baseline_profile-min.json
var rev5LowBaseline []byte

// Rev5CatalogWrapper embeds oscalTypes.Catalog for extension or additional methods.
type rev5CatalogWrapper struct {
	Catalog oscalTypes.Catalog `json:"catalog"`
}

func Rev5Catalog() *Catalog          { return MustParseCatalogJSON(rev5Catalog) }
func Rev5HighBaseline() *Catalog     { return MustParseCatalogJSON(rev5HighBaseline) }
func Rev5ModerateBaseline() *Catalog { return MustParseCatalogJSON(rev5ModerateBaseline) }
func Rev5LowBaseline() *Catalog      { return MustParseCatalogJSON(rev5LowBaseline) }

func MustParseCatalogJSON(b []byte) *Catalog {
	w := rev5CatalogWrapper{}
	err := json.Unmarshal(b, &w)
	if err != nil {
		panic(err)
	}
	c := Catalog(w.Catalog)
	return &c
}
