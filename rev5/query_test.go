package rev5

import (
	"slices"
	"strings"
	"testing"
)

func TestQueryBaselines(t *testing.T) {
	c := CatalogAll()
	ids := ControlIDs{}
	want := map[string]int{
		BaselineLow:      len(ids.LowBaseline()),
		BaselineModerate: len(ids.ModerateBaseline()),
		BaselineHigh:     len(ids.HighBaseline()),
	}
	for b, n := range want {
		got, err := c.Query(Query{Baseline: b})
		if err != nil {
			t.Fatal(err)
		}
		if len(got) != n {
			t.Errorf("Query(%s) = %d controls, want %d", b, len(got), n)
		}
	}
}

func TestQueryFilters(t *testing.T) {
	c := CatalogAll()
	all, err := c.Query(Query{IncludeWithdrawn: true})
	if err != nil {
		t.Fatal(err)
	}
	active, err := c.Query(Query{})
	if err != nil {
		t.Fatal(err)
	}
	if len(active) >= len(all) {
		t.Errorf("withdrawn controls not left out: %d active of %d", len(active), len(all))
	}
	for _, i := range active {
		if i.Withdrawn {
			t.Fatalf("%s is withdrawn but was listed", i.ID)
		}
	}

	sys, err := c.Query(Query{Baseline: BaselineLow, Families: []string{"AC"}, Level: PropImplementationLevelValueSystem})
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, i := range sys {
		if i.Family != "AC" || !slices.Contains(i.ImplementationLevels, PropImplementationLevelValueSystem) {
			t.Errorf("%s does not match the filter: %+v", i.ID, i)
		}
		got = append(got, i.ID)
	}
	if !slices.Contains(got, "AC-3") || slices.Contains(got, "AC-1") {
		t.Errorf("low AC system controls = %v, want AC-3 and not AC-1", got)
	}
}

func TestQueryValidate(t *testing.T) {
	c := CatalogAll()
	for _, q := range []Query{
		{Baseline: "extreme"},
		{Families: []string{"zz"}},
		{Level: "team"},
	} {
		if _, err := c.Query(q); err == nil {
			t.Errorf("Query(%+v) succeeded, want an error", q)
		}
	}
}

func TestControlInfo(t *testing.T) {
	c := CatalogAll()
	info, err := c.ControlInfo("ac-2.1")
	if err != nil {
		t.Fatal(err)
	}
	if info.ID != "AC-2 (1)" || info.OSCALID != "ac-2.1" || info.Family != "AC" || !info.Enhancement {
		t.Errorf("unexpected identity: %+v", info)
	}
	if !slices.Equal(info.Baselines, []string{BaselineModerate, BaselineHigh}) {
		t.Errorf("Baselines = %v, want moderate, high", info.Baselines)
	}
	if info.Statement == "" || info.Guidance == "" {
		t.Error("statement and guidance should be set")
	}

	base, err := c.ControlInfo("AC-2")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(base.Statement, "a. ") || !strings.Contains(base.Statement, "\n  1. ") {
		t.Errorf("statement should be a labelled outline, got:\n%s", base.Statement)
	}

	if _, err := c.ControlInfo("AC-99"); err == nil {
		t.Error("ControlInfo(AC-99) succeeded, want an error")
	}
}

func TestCheckIDs(t *testing.T) {
	got := CatalogAll().CheckIDs([]string{"AC-2", "ac-2.1", "AC-13", "AC-99", "XX-1"})
	want := []string{IDCheckOK, IDCheckOK, IDCheckWithdrawn, IDCheckUnknown, IDCheckInvalid}
	for i, r := range got {
		if r.Status != want[i] {
			t.Errorf("CheckIDs(%s) = %s, want %s", r.Input, r.Status, want[i])
		}
	}
	if got[1].ID != "AC-2 (1)" {
		t.Errorf("ID = %q, want AC-2 (1)", got[1].ID)
	}
}

func TestFamilyInfos(t *testing.T) {
	fams, err := CatalogAll().FamilyInfos()
	if err != nil {
		t.Fatal(err)
	}
	if len(fams) != 20 {
		t.Errorf("got %d families, want 20", len(fams))
	}
	low := 0
	for _, f := range fams {
		low += f.Baselines[BaselineLow]
	}
	if n := len(ControlIDs{}.LowBaseline()); low != n {
		t.Errorf("family low counts sum to %d, want %d", low, n)
	}
}
