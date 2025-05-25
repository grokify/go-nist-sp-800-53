package rev5

import "testing"

func TestControlSetHighUplift(t *testing.T) {
	ids := ControlIDs{}
	testFn := func(t *testing.T, cs *ControlSet, tierName string) {
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
	testFn(t, ControlSetHighBaseline(), TierHighBaseline)
	testFn(t, ControlSetHighUplift(), TierHighUplift)
	testFn(t, ControlSetModerateBaseline(), TierModerateBaseline)
	testFn(t, ControlSetModerateUplift(), TierModerateUplift)
	testFn(t, ControlSetLowBaseline(), TierLowBaseline)
}
