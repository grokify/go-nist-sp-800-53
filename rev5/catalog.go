package rev5

import (
	"fmt"
	"sort"
	"strings"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

type Catalog oscalTypes.Catalog

type IDType string

const (
	IDTypeOSCAL     IDType = "id"
	IDTypeOSCALSort IDType = "sort-id"
	IDTypeNIST      IDType = "label"
	IDTypeNISTSort  IDType = "sp800-53a"
)

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
								switch idtype {
								case IDTypeOSCAL:
									out = append(out, control.ID)
								default:
									for _, prop := range *control.Props {
										if idtype == IDTypeOSCALSort {
											if prop.Name == string(idtype) {
												if v := strings.TrimSpace(prop.Value); v != "" {
													out = append(out, v)
												}
											}
										} else if idtype == IDTypeNISTSort {
											if prop.Name == PropNameLabel && prop.Class == string(idtype) {
												if v := strings.TrimSpace(prop.Value); v != "" {
													out = append(out, v)
												}
											}
										} else if idtype == IDTypeNIST {
											if prop.Name == PropNameLabel && prop.Class == "" {
												if v := strings.TrimSpace(prop.Value); v != "" {
													out = append(out, v)
												}
											}
										} else {
											return []string{}, fmt.Errorf("idtype not supported (%s)", idtype)
										}
									}
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
