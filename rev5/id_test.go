package rev5

import "testing"

var idTests = []struct {
	v             string
	wantNIST      string
	wantNISTSort  string
	wantOSCAL     string
	wantOSCALSort string
}{
	{"ac-1", "AC-1", "AC-01", "ac-1", "ac-01"},
	{"ac-1.1", "AC-1 (1)", "AC-01 (01)", "ac-1.1", "ac-01.01"},
	{"ac-1.2", "AC-1 (2)", "AC-01 (02)", "ac-1.2", "ac-01.02"},
	{"pe-3.1", "PE-3 (1)", "PE-03 (01)", "pe-3", "pe-03"},
}

func TestID(t *testing.T) {
	for _, tt := range idTests {
		try, err := ParseIDFromOSCAL(tt.v)
		if err != nil {
			t.Errorf("rev5.ParseIDFromOSCAL(\"%s\") error: (%s)", tt.v, err.Error())
		} else {
			tryNIST, err := try.FormatNIST()
			if err != nil {
				t.Errorf("rev5.FormatNIST() error: on (%s) error (%s)", tt.v, err.Error())
			} else if tryNIST != tt.wantNIST {
				t.Errorf("rev5.FormatNIST() mimatch: on (%s) want (%s) got (%s)", tt.v, tt.wantNIST, tryNIST)
			}
		}
	}
}
