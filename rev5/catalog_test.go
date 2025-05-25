package rev5

import (
	"fmt"
	"testing"
)

func TestCatalogControls(t *testing.T) {
	cat := Rev5Catalog()
	ids := ControlIDs{}
	testCatalogControls(t, cat, ids.HighBaseline())
	testCatalogControls(t, cat, ids.HighUplift())
	testCatalogControls(t, cat, ids.ModerateBaseline())
	testCatalogControls(t, cat, ids.ModerateUplift())
	testCatalogControls(t, cat, ids.LowBaseline())
}

func testCatalogControls(t *testing.T, cat *Catalog, ctrIDs []string) {
	fmt.Printf("TESTING (%d)\n", len(ctrIDs))
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
