package rev5

import (
	"slices"
	"testing"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

func TestImplementationLevels(t *testing.T) {
	cs := ControlSetLowBaseline()
	tests := []struct {
		id   string
		want []string
	}{
		{"ac-1", []string{PropImplementationLevelValueOrganization}},
		{"ac-3", []string{PropImplementationLevelValueSystem}},
		{"sc-7", []string{PropImplementationLevelValueSystem}},
	}
	for _, tt := range tests {
		ctr, ok := cs.Get(tt.id)
		if !ok {
			t.Fatalf("control %s not found in low baseline", tt.id)
		}
		if got := ctr.ImplementationLevels(); !slices.Equal(got, tt.want) {
			t.Errorf("%s ImplementationLevels() = %v, want %v", tt.id, got, tt.want)
		}
		wantOrg := slices.Contains(tt.want, PropImplementationLevelValueOrganization)
		if got := ctr.HasImplementationLevelOrganization(); got != wantOrg {
			t.Errorf("%s HasImplementationLevelOrganization() = %v, want %v", tt.id, got, wantOrg)
		}
		wantSys := slices.Contains(tt.want, PropImplementationLevelValueSystem)
		if got := ctr.HasImplementationLevelSystem(); got != wantSys {
			t.Errorf("%s HasImplementationLevelSystem() = %v, want %v", tt.id, got, wantSys)
		}
	}
}

func TestImplementationLevelsDedupe(t *testing.T) {
	props := []oscalTypes.Property{
		{Name: PropNameImplementationLevel, Value: "system"},
		{Name: PropNameImplementationLevel, Value: " Organization "},
		{Name: PropNameImplementationLevel, Value: "system"},
		{Name: PropNameImplementationLevel, Value: ""},
		{Name: PropNameLabel, Value: "AC-1"},
	}
	ctr := Control{Props: &props}
	want := []string{"organization", "system"}
	if got := ctr.ImplementationLevels(); !slices.Equal(got, want) {
		t.Errorf("ImplementationLevels() = %v, want %v", got, want)
	}
	if got := (Control{}).ImplementationLevels(); len(got) != 0 {
		t.Errorf("ImplementationLevels() without props = %v, want none", got)
	}
}

func TestFilterImplementationLevel(t *testing.T) {
	cs := ControlSetLowBaseline()
	org, err := cs.FilterImplementationLevel(true, false)
	if err != nil {
		t.Fatal(err)
	}
	sys, err := cs.FilterImplementationLevel(false, true)
	if err != nil {
		t.Fatal(err)
	}
	both, err := cs.FilterImplementationLevel(true, true)
	if err != nil {
		t.Fatal(err)
	}
	none, err := cs.FilterImplementationLevel(false, false)
	if err != nil {
		t.Fatal(err)
	}
	if org.Len() == 0 || sys.Len() == 0 {
		t.Fatalf("expected organization and system controls, got %d and %d", org.Len(), sys.Len())
	}
	if both.Len() != cs.Len() {
		t.Errorf("both levels = %d controls, want all %d", both.Len(), cs.Len())
	}
	if none.Len() != 0 {
		t.Errorf("no levels = %d controls, want 0", none.Len())
	}
	if org.Len()+sys.Len() < both.Len() {
		t.Errorf("organization (%d) + system (%d) < both (%d)", org.Len(), sys.Len(), both.Len())
	}
	if _, ok := org.Get("ac-1"); !ok {
		t.Error("ac-1 missing from organization-level controls")
	}
	if _, ok := sys.Get("ac-3"); !ok {
		t.Error("ac-3 missing from system-level controls")
	}
}

func TestControlsIDsImplementationLevel(t *testing.T) {
	cat := CatalogLowBaseline()
	var ac Controls
	for _, g := range *cat.Groups {
		if g.ID == "ac" {
			ac = Controls(*g.Controls)
		}
	}
	if len(ac) == 0 {
		t.Fatal("ac family not found in low baseline")
	}
	org, err := ac.IDsImplementationLevelOrganization()
	if err != nil {
		t.Fatal(err)
	}
	sys, err := ac.IDsImplementationLevelSystem()
	if err != nil {
		t.Fatal(err)
	}
	has := func(ids IDs, s string) bool {
		for _, id := range ids {
			if f, err := id.FormatNIST(); err == nil && f == s {
				return true
			}
		}
		return false
	}
	if !has(org, "AC-1") || has(org, "AC-3") {
		t.Errorf("organization IDs should include AC-1 and not AC-3: %v", org)
	}
	if !has(sys, "AC-3") || has(sys, "AC-1") {
		t.Errorf("system IDs should include AC-3 and not AC-1: %v", sys)
	}
}
