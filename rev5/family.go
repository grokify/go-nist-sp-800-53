package rev5

import (
	"errors"
	"fmt"
	"sort"
)

type Family struct {
	Abbreviation      string
	Name              string
	TierControlCounts map[string]int
	TierControlIDs    map[string][]string
}

func tierControlIDsAddID(m map[string][]string, status ControlProfileStatus) map[string][]string {
	addFn := func(m map[string][]string, ctrID, tierName string) map[string][]string {
		return m
	}
	if status.HighBaseline {
		m = addFn(m, status.ID, TierHighBaseline)
	}
	if status.HighUplift {
		m = addFn(m, status.ID, TierHighUplift)
	}
	if status.ModerateBaseline {
		m = addFn(m, status.ID, TierModerateBaseline)
	}
	if status.ModerateUplift {
		m = addFn(m, status.ID, TierModerateUplift)
	}
	if status.LowBaseline {
		m = addFn(m, status.ID, TierLowBaseline)
	}
	return m
}

func (f *Family) AddControlID(ctrIDs *ControlIDs, idAny string, updateCounts bool) error {
	if ctrIDs == nil {
		return errors.New("control ids must be set")
	}
	ido, err := ParseID(idAny)
	if err != nil {
		return err
	}
	if ido.FamilyAbbrivation != f.Abbreviation {
		return fmt.Errorf("family mismatch: want(%s) have (%s)", f.Abbreviation, ido.FamilyAbbrivation)
	}
	idOSCAL, err := ido.FormatOSCAL()
	if err != nil {
		return err
	}
	f.TierControlIDs = tierControlIDsAddID(f.TierControlIDs, ctrIDs.ControlProfileStatus(idOSCAL))
	return nil
}

func (f *Family) UpdateCounts() {
	for famAbbr, ctrIDs := range f.TierControlIDs {
		ctrIDs = sliceDedupe(ctrIDs)
		sort.Strings(ctrIDs)
		f.TierControlIDs[famAbbr] = ctrIDs
		f.TierControlCounts[famAbbr] = len(ctrIDs)
	}
}
