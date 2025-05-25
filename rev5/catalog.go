package rev5

import (
	"errors"
	"fmt"
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

func (c *Catalog) Control(controlID string) (*Control, error) {
	controlID = strings.TrimSpace(controlID)
	if controlID == "" {
		return nil, errors.New("controlid cannot be empty)")
	}
	if c.Controls != nil {
		for _, ctr := range *c.Controls {
			if ctr.ID == controlID {
				c2 := Control(ctr)
				return &c2, nil
			}
		}
	}
	if c.Groups != nil {
		for _, grp := range *c.Groups {
			if grp.Controls != nil {
				for _, ctr := range *grp.Controls {
					if ctr.ID == controlID {
						c2 := Control(ctr)
						return &c2, nil
					}
					// Control appears most likely to be here, though check above for
					// consistency.
					if ctr.Controls != nil {
						for _, ctr2 := range *ctr.Controls {
							if ctr2.ID == controlID {
								c2 := Control(ctr2)
								return &c2, nil
							}
						}
					}
				}
			}
		}
	}
	return nil, fmt.Errorf("control not found for id (%s)", controlID)
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
