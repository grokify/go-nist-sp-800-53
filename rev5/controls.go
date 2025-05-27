package rev5

import (
	"strings"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

type Controls []oscalTypes.Control

// Filter filters on IDTypeOSCAL
func (ctrs Controls) Filter(ctlIDsOSCAL []string) Controls {
	var out Controls
	incl := map[string]int{}
	for _, id := range ctlIDsOSCAL {
		if id = strings.TrimSpace(id); id != "" {
			incl[id]++
		}
	}
	if len(incl) == 0 {
		return out
	}
	for _, ctr := range ctrs {
		if len(incl) > 0 {
			if _, ok := incl[ctr.ID]; !ok {
				continue
			}
		}
		out = append(out, ctr)
	}
	return out
}

func (ctrs Controls) Flatten() Controls {
	var out Controls
	for _, ctr := range ctrs {
		out = append(out, ctr)
		if ctr.Controls != nil {
			subCtrs := Controls(*ctr.Controls)
			subCtrsFlatten := subCtrs.Flatten()
			out = append(out, subCtrsFlatten...)
		}
	}
	return out
}
