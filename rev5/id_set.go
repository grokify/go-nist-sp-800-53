package rev5

// IDSet is a set of IDs. The key is OSCAL.
type IDSet struct {
	Data map[string]ID
}

func NewIDSet() *IDSet {
	return &IDSet{Data: map[string]ID{}}
}

func (set *IDSet) Add(ids ...ID) error {
	for _, id := range ids {
		if err := id.idPartsVerify(); err != nil {
			return err
		} else if idOSCAL, err := id.FormatOSCAL(); err != nil {
			return err
		} else {
			set.Data[idOSCAL] = id
			return nil
		}
	}
	return nil
}

func (set *IDSet) AddIDString(s ...string) error {
	for _, si := range s {
		if id, err := ParseID(si); err != nil {
			return err
		} else if err := set.Add(id); err != nil {
			return err
		}
	}
	return nil
}

func (set *IDSet) FamilyCounts() map[string]int {
	out := map[string]int{}
	for _, id := range set.Data {
		out[id.FamilyAbbreviation]++
	}
	return out
}

func (set *IDSet) FilterBaseControls(inclEnhancementBase bool) (*IDSet, error) {
	out := NewIDSet()
	for _, id := range set.Data {
		if id.EnhancementNumber <= 0 {
			if err := out.Add(id); err != nil {
				return nil, err
			}
		} else if idBase, err := id.IDBase(); err != nil {
			return nil, err
		} else if err := out.Add(idBase); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (set *IDSet) IDs() IDs {
	var out []ID
	for _, id := range set.Data {
		out = append(out, id)
	}
	return out
}
