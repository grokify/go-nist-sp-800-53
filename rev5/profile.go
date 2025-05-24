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
