package rev5

import (
	"strings"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

type Catalog oscalTypes.Catalog

func (c *Catalog) Families(full bool) []oscalTypes.Group {
	var out []oscalTypes.Group
	for _, g1 := range *c.Groups {
		if g1.Class != GroupClassFamily {
			continue
		}
		if full {
			out = append(out, g1)
		} else {
			out = append(out, oscalTypes.Group{
				ID:    g1.ID,
				Class: g1.Class,
				Title: g1.Title,
			})
		}
	}
	return out
}

func (c *Catalog) FamiliesMap() map[string]string {
	fams := c.Families(false)
	out := map[string]string{}
	for _, f := range fams {
		out[strings.ToUpper(f.ID)] = f.Title
	}
	return out
}
