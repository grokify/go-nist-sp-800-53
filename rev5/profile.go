package rev5

import (
	"sort"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

type Profile oscalTypes.Profile

func (p *Profile) ControlIDs() []string {
	var out []string
	for _, import1 := range p.Imports {
		for _, includeControl := range *import1.IncludeControls {
			out = append(out, *includeControl.WithIds...)
		}
	}
	sort.Strings(out)
	return out
}

func (p *Profile) ControlIDsMap() map[string]string {
	out := map[string]string{}
	for _, id := range p.ControlIDs() {
		out[id] = id
	}
	return out
}

func (p *Profile) ControlSet(cat *Catalog) (*ControlSet, error) {
	cs := ControlSet{}
	for _, id := range p.ControlIDs() {
		if ctr, err := cat.Control(id); err != nil {
			return nil, err
		} else if err := cs.Add(*ctr); err != nil {
			return nil, err
		}
	}
	return &cs, nil
}

/*

func (p *Profile) ControlIDStatusCounts(cat *Catalog) (map[string]int, error) {
	if cat == nil {
		return map[string]int{}, errors.New("catalog cannot be nil")
	}
	ids := p.ControlIDs()
	return cat.ControlIDStatusCounts(ids), nil
}
*/
