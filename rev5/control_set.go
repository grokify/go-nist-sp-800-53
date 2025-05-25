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

func (cs *ControlSet) AddAll(ctrls []Control) {
	for _, ctr := range ctrls {
		cs.Add(ctr)
	}
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
