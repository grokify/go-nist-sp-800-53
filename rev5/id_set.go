package rev5

// IDSet is a set of IDs. The key is OSCAL.
type IDSet struct {
	Data map[string]ID
}

func (set *IDSet) AddID(ids ...ID) error {
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
		} else if err := set.AddID(id); err != nil {
			return err
		}
	}
	return nil
}

func (set *IDSet) FamilyCounts() map[string]int {
	out := map[string]int{}
	for _, id := range set.Data {
		out[id.FamilyAbbrivation]++
	}
	return out
}
