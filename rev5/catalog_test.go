package rev5

import (
	"testing"
)

func TestCatalogControls(t *testing.T) {
	cat := Rev5Catalog()
	ids := ControlIDs{}
	for _, tierName := range Tiers() {
		if tierIDs, err := ids.Tier(tierName); err != nil {
			t.Errorf("tier not found(\"%s\") error: (%s)", tierName, err.Error())
		} else {
			testCatalogControls(t, cat, tierName, tierIDs)
		}
	}
}

func testCatalogControls(t *testing.T, cat *Catalog, label string, ctrIDs []string) {
	t.Logf("testing existence of (%d) (%s) controls", len(ctrIDs), label)
	for _, ctrID := range ctrIDs {
		ctr, err := cat.Control(ctrID)
		if err != nil {
			t.Errorf("high uplift control not found(\"%s\") error: (%s)", ctrID, err.Error())
		} else if ctr == nil {
			t.Errorf("high uplift control found but is nil(\"%s\")", ctrID)
		} else if ctr.ID != ctrID {
			t.Errorf("high uplift control found but has mismatch(\"%s\") found (%s)", ctrID, ctr.ID)
		}
	}
}
