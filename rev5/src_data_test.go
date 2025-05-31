package rev5

import "testing"

func TestTierCatalogControls(t *testing.T) {
	cat := CatalogAll()
	ids := ControlIDs{}
	testFnCatalogControls := func(t *testing.T, cat *Catalog, label string, ctrIDs []string) {
		t.Logf("testing existence of (%d) (%s) controls", len(ctrIDs), label)
		for _, ctrID := range ctrIDs {
			ctr, err := cat.ControlByIDStringAny(ctrID)
			if err != nil {
				t.Errorf("high uplift control not found(\"%s\") error: (%s)", ctrID, err.Error())
			} else if ctr == nil {
				t.Errorf("high uplift control found but is nil(\"%s\")", ctrID)
			} else if ctr.ID != ctrID {
				t.Errorf("high uplift control found but has mismatch(\"%s\") found (%s)", ctrID, ctr.ID)
			}
		}
	}
	for _, tierName := range Tiers() {
		if tierIDs, err := ids.Tier(tierName); err != nil {
			t.Errorf("tier not found(\"%s\") error: (%s)", tierName, err.Error())
		} else {
			testFnCatalogControls(t, cat, tierName, tierIDs)
		}
	}
}

func TestTierControlSets(t *testing.T) {
	ids := ControlIDs{}
	testFnControlSet := func(t *testing.T, cs *ControlSet, tierName string) {
		if cs == nil {
			t.Fatalf("no control set found for tier (%s)", tierName)
		} else if cs.Len() == 0 {
			t.Fatalf("control set count mismatch for tier (%s)", tierName)
		} else {
			var tierIDs []string
			switch tierName {
			case TierHighBaseline:
				tierIDs = ids.HighBaseline()
			case TierHighUplift:
				tierIDs = ids.HighUplift()
			case TierModerateBaseline:
				tierIDs = ids.ModerateBaseline()
			case TierModerateUplift:
				tierIDs = ids.ModerateUplift()
			case TierLowBaseline:
				tierIDs = ids.LowBaseline()
			}
			if cs.Len() != len(tierIDs) {
				t.Fatalf("control set count mismatch ids (%d) controls (%d) for tier (%s)", len(tierIDs), cs.Len(), tierName)
			}
		}
	}
	testFnControlSet(t, ControlSetHighBaseline(), TierHighBaseline)
	testFnControlSet(t, ControlSetHighUplift(), TierHighUplift)
	testFnControlSet(t, ControlSetModerateBaseline(), TierModerateBaseline)
	testFnControlSet(t, ControlSetModerateUplift(), TierModerateUplift)
	testFnControlSet(t, ControlSetLowBaseline(), TierLowBaseline)
}
