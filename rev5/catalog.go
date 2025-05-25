package rev5

import (
	"sort"
	"strings"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

type Catalog oscalTypes.Catalog

func (c *Catalog) MustEnhancementIDs(idtype IDType, emptyOnError bool) []string {
	if ids, err := c.EnhancementIDs(idtype); err == nil {
		return ids
	} else if emptyOnError {
		return []string{}
	} else {
		panic(err)
	}
}

func (c *Catalog) EnhancementIDs(idtype IDType) ([]string, error) {
	if strings.TrimSpace(string(idtype)) == "" {
		idtype = IDTypeOSCAL
	}
	var out []string
	if c.Groups != nil {
		for _, group := range *c.Groups {
			if group.Controls != nil {
				for _, controlw := range *group.Controls {
					if controlw.Controls != nil {
						for _, control := range *controlw.Controls {
							if control.Class == ClassSP80053Enhancement {
								c2 := Control(control)
								if idval, err := c2.IDForType(idtype); err != nil {
									return []string{}, err
								} else {
									out = append(out, idval)
								}
							}
						}
					}
				}
			}
		}
	}
	sort.Strings(out)
	return out, nil
}

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
