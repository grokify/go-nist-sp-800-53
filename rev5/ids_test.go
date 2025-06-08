package rev5

import (
	"slices"
	"strings"
	"testing"
)

type idsTest struct {
	v             []string
	wantNIST      []string
	wantNISTSort  []string
	wantOSCAL     []string
	wantOSCALSort []string
}

var idsTests = []idsTest{
	{
		[]string{"AC-10", "AC-18 (4)", "AC-18 (5)", "AC-2 (11)", "AC-2 (12)", "AC-4 (4)", "AC-6 (3)"},
		[]string{"AC-2 (11)", "AC-2 (12)", "AC-4 (4)", "AC-6 (3)", "AC-10", "AC-18 (4)", "AC-18 (5)"},
		[]string{"AC-02 (11)", "AC-02 (12)", "AC-04 (04)", "AC-06 (03)", "AC-10", "AC-18 (04)", "AC-18 (05)"},
		[]string{"ac-2.11", "ac-2.12", "ac-4.4", "ac-6.3", "ac-10", "ac-18.4", "ac-18.5"},
		[]string{"ac-02.11", "ac-02.12", "ac-04.04", "ac-06.03", "ac-10", "ac-18.04", "ac-18.05"},
	},
}

func TestIDsParseSort(t *testing.T) {
	for _, tt := range idsTests {
		ids, err := ParseIDs(tt.v)
		if err != nil {
			t.Errorf("rev5.ParseIDs(%v) error: (%s)", tt.v, err.Error())
			continue
		}
		tryIDsNIST, err := ids.FormatNIST()
		if err != nil {
			t.Errorf("ids.FormatNIST() error: (%s)", err.Error())
		} else if !slices.Equal(tt.wantNIST, tryIDsNIST) {
			t.Errorf("ids.FormatNIST() mismatch: want (%s) got (%s)",
				strings.Join(tt.wantNIST, ","), strings.Join(tryIDsNIST, ","))
		}
		tryIDsNISTSort, err := ids.FormatNISTSort()
		if err != nil {
			t.Errorf("ids.FormatNISTSort() error: (%s)", err.Error())
		} else if !slices.Equal(tt.wantNISTSort, tryIDsNISTSort) {
			t.Errorf("ids.FormatNISTSort() mismatch: want (%s) got (%s)",
				strings.Join(tt.wantNISTSort, ","), strings.Join(tryIDsNISTSort, ","))
		}
		tryIDsOSCAL, err := ids.FormatOSCAL()
		if err != nil {
			t.Errorf("ids.FormatOSCALSort() error: (%s)", err.Error())
		} else if !slices.Equal(tt.wantOSCAL, tryIDsOSCAL) {
			t.Errorf("ids.FormatOSCALSort() mismatch: want (%s) got (%s)",
				strings.Join(tt.wantOSCAL, ","), strings.Join(tryIDsOSCAL, ","))
		}
		tryIDsOSCALSort, err := ids.FormatOSCALSort()
		if err != nil {
			t.Errorf("ids.FormatOSCALSort() error: (%s)", err.Error())
		} else if !slices.Equal(tt.wantOSCALSort, tryIDsOSCALSort) {
			t.Errorf("ids.FormatOSCALSort() mismatch: want (%s) got (%s)",
				strings.Join(tt.wantOSCALSort, ","), strings.Join(tryIDsOSCALSort, ","))
		}
	}
}
