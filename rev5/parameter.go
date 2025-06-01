package rev5

import (
	"fmt"
	"slices"
	"strings"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

type Parameter oscalTypes.Parameter

func (p Parameter) AltIDs() []string {
	var out []string
	if p.Props != nil {
		for _, prop := range *p.Props {
			if prop.Name == PropNameAltIdentifier {
				if paramAltID := strings.TrimSpace(prop.Value); paramAltID != "" {
					out = append(out, paramAltID)
				}
			}
		}
	}
	return out
}

func (p Parameter) ProseLines() []string {
	var out []string
	if p.Guidelines != nil {
		for _, gl := range *p.Guidelines {
			out = append(out, gl.Prose)
		}
	}
	return out
}

func (p Parameter) ProseString(sep string) string {
	return strings.TrimSpace(strings.Join(p.ProseLines(), sep))
}

func (p Parameter) SelectString() string {
	if p.Select == nil {
		return ""
	} else if p.Select.Choice == nil {
		return ""
	}
	var conj string
	switch p.Select.HowMany {
	case "":
		conj = "or"
	case ParamHowManyOneOrMore:
		conj = "and/or"
	default:
		fmt.Printf("%v\n", p)
		panic(fmt.Sprintf("parameter.SelectString(): unknown how many (%s)", p.Select.HowMany))
	}

	parts := slices.Clone(*p.Select.Choice)
	if len(parts) == 0 {
		return ""
	} else if len(parts) == 1 {
		return strings.TrimSpace(parts[0])
	} else if len(parts) == 2 {
		return strings.TrimSpace(strings.Join(parts, conj))
	} else {
		if conj != "" {
			parts[len(parts)-1] = conj + " " + parts[len(parts)-1]
		}
		return strings.TrimSpace(strings.Join(parts, ", "))
	}
}

func (p Parameter) ValueString() string {
	if v := p.ProseString(SepLines); v != "" {
		return v
	} else if v := p.SelectString(); v != "" {
		return v
	} else if v := strings.TrimSpace(p.Label); v != "" {
		return v
	} else {
		fmt.Printf("%v\n", p)
		panic("parameter display value string not found")
	}
}
