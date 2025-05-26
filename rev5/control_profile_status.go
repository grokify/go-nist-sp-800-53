package rev5

type ControlProfileStatusSet struct {
	Data map[string]ControlProfileStatus // key = OSCAL ID
}

type ControlProfileStatus struct {
	ID               string // OSCAL ID
	HighBaseline     bool
	HighUplift       bool
	ModerateBaseline bool
	ModerateUplift   bool
	LowBaseline      bool
}

func (set *ControlProfileStatusSet) AddIDsOSCAL(ids []string) {
	profileIDs := ControlIDs{}
	for _, id := range ids {
		set.Data[id] = profileIDs.ControlProfileStatus(id)
	}
}
