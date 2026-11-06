package rev5

import (
	"errors"
	"strings"
)

type ControlSet struct {
	Data map[string]Control
}

func NewControlSet() *ControlSet {
	return &ControlSet{Data: map[string]Control{}}
}

func (cs *ControlSet) Add(ctr Control) error {
	if cs.Data == nil {
		cs.Data = map[string]Control{}
	}
	if strings.TrimSpace(ctr.ID) == "" {
		return errors.New("control id cannot be empty")
	} else {
		cs.Data[ctr.ID] = ctr
		return nil
	}
}

func (cs *ControlSet) AddAll(ctrls []Control) error {
	for _, ctr := range ctrls {
		if err := cs.Add(ctr); err != nil {
			return err
		}
	}
	return nil
}

func (cs *ControlSet) FamilySet(addZeros bool) (*FamilySet, error) {
	set := NewFamilySet(false)
	for ctrID := range cs.Data {
		if err := set.AddControlID(ctrID, false); err != nil {
			return nil, err
		}
	}
	set.UpdateCounts()

	return set, nil
}

// FilterImplementationLevel returns a new ControlSet containing the controls
// whose implementation level matches one of the requested levels. A control is
// included when inclOrganization is true and it has organization-level
// implementation, or inclSystem is true and it has system-level implementation.
// If both flags are false, the result is empty. It returns an error if a
// control cannot be added (e.g. an empty control ID), matching the
// error-returning style of FamilySet and IDsSet.
func (cs *ControlSet) FilterImplementationLevel(inclOrganization, inclSystem bool) (*ControlSet, error) {
	out := NewControlSet()
	for _, ctr := range cs.Data {
		if (inclOrganization && ctr.HasImplementationLevelOrganization()) ||
			(inclSystem && ctr.HasImplementationLevelSystem()) {
			if err := out.Add(ctr); err != nil {
				return nil, err
			}
		}
	}
	return out, nil
}

func (cs *ControlSet) Get(id string) (Control, bool) {
	ctr, ok := cs.Data[id]
	return ctr, ok
}

func (cs *ControlSet) GetAll() []Control {
	ctrls := []Control{}
	for _, ctr := range cs.Data {
		ctrls = append(ctrls, ctr)
	}
	return ctrls
}

func (cs *ControlSet) IDsSet() (*IDSet, error) {
	is := NewIDSet()
	for _, ctr := range cs.Data {
		if id, err := ctr.ControlID(); err != nil {
			return nil, err
		} else if err := is.Add(id); err != nil {
			return nil, err
		}
	}
	return is, nil
}

func (cs *ControlSet) Len() int {
	return len(cs.Data)
}
