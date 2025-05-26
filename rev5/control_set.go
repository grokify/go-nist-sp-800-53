package rev5

import (
	"errors"
	"strings"
)

type ControlSet struct {
	Data map[string]Control
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

func (cs *ControlSet) Len() int {
	return len(cs.Data)
}
