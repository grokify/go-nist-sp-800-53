package rev5

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
	"github.com/grokify/go-nist-sp-800-53/rev5/rev5tmpl"
)

type ParameterSet struct {
	Data         map[string]Parameter
	TemplateVars [][]string
}

func NewParameterSet() *ParameterSet {
	return &ParameterSet{Data: map[string]Parameter{}}
}

func ParameterSetFromControl(ctr *Control) (*ParameterSet, error) {
	if ctr == nil {
		return nil, errors.New("control cannot be nil")
	} else {
		set := NewParameterSet()
		return set, set.AddParameters(ctr.Params)
	}
}

func (set *ParameterSet) AddParameters(params *[]oscalTypes.Parameter) error {
	if params == nil {
		return nil
	}
	for _, param := range *params {
		if err := set.Add(Parameter(param)); err != nil {
			return err
		}
	}
	return nil
}

func (set *ParameterSet) Add(param Parameter) error {
	if strings.TrimSpace(param.ID) == "" {
		return errors.New("parameter id cannot be nil or empty")
	}
	if set.Data == nil {
		set.Data = map[string]Parameter{}
	}
	if err := set.addByID(param.ID, param); err != nil {
		return err
	}
	if altIDs := param.AltIDs(); len(altIDs) > 0 {
		for _, altID := range altIDs {
			if err := set.addByID(altID, param); err != nil {
				return err
			}
		}
	}
	return nil
}

func (set *ParameterSet) addByID(paramID string, param Parameter) error {
	sep := "\n\n"
	addProseString := param.ProseString(sep)
	if try := set.Get(paramID); try != nil {
		if existProseString := try.ProseString(sep); addProseString != existProseString {
			return fmt.Errorf("prose mismatch, existing (%s) adding (%s)", existProseString, addProseString)
		} else {
			return nil
		}
	}
	set.Data[paramID] = param
	return nil
}

func (set *ParameterSet) Get(paramID string) *Parameter {
	if set.Data == nil {
		set.Data = map[string]Parameter{}
	}
	if param, ok := set.Data[paramID]; ok {
		return &param
	} else {
		return nil
	}
}

func (set *ParameterSet) Keys() []string {
	var out []string
	for k := range set.Data {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func (set *ParameterSet) MapKeyValue() map[string]string {
	out := map[string]string{}
	for k, param := range set.Data {
		if v := param.ValueString(); v == "" {
			panic("no value")
		} else {
			out[k] = v
		}
	}
	return out
}

func (set *ParameterSet) ProseForParamID(paramID string) ([]string, error) {
	var out []string
	param, ok := set.Data[paramID]
	if !ok {
		return []string{}, errors.New("paramID not found")
	}
	if param.Guidelines != nil {
		for _, gl := range *param.Guidelines {
			if p := strings.TrimSpace(gl.Prose); p != "" {
				out = append(out, p)
			}
		}
	}
	return out, nil
}

// RenderTemplate renders a template string with the parameters in the
// parameter set.
func (set *ParameterSet) RenderTemplate(template string) (string, error) {
	return rev5tmpl.Render(template, set.MapKeyValue())
}
