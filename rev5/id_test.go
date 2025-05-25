package rev5

import "testing"

var idTests = []struct {
	vType         IDType
	v             string
	wantNIST      string
	wantNISTSort  string
	wantOSCAL     string
	wantOSCALSort string
}{
	{IDTypeOSCAL, "ac-1", "AC-1", "AC-01", "ac-1", "ac-01"},
	{IDTypeOSCAL, "AC-1", "AC-1", "AC-01", "ac-1", "ac-01"},
	{IDTypeOSCAL, "ac-1.1", "AC-1 (1)", "AC-01 (01)", "ac-1.1", "ac-01.01"},
	{IDTypeOSCAL, "ac-1.2", "AC-1 (2)", "AC-01 (02)", "ac-1.2", "ac-01.02"},
	{IDTypeOSCAL, "pe-3", "PE-3", "PE-03", "pe-3", "pe-03"},
	{IDTypeOSCAL, "pe-3.1", "PE-3 (1)", "PE-03 (01)", "pe-3.1", "pe-03.01"},

	{IDTypeNIST, "AC-1", "AC-1", "AC-01", "ac-1", "ac-01"},
	{IDTypeNIST, "ac-1(1)", "AC-1 (1)", "AC-01 (01)", "ac-1.1", "ac-01.01"},
	{IDTypeNIST, "AC-1(1)", "AC-1 (1)", "AC-01 (01)", "ac-1.1", "ac-01.01"},
	{IDTypeNIST, "AC-1(2)", "AC-1 (2)", "AC-01 (02)", "ac-1.2", "ac-01.02"},
	{IDTypeNIST, "AC-01(02)", "AC-1 (2)", "AC-01 (02)", "ac-1.2", "ac-01.02"},
	{IDTypeNIST, "AC-1 (2)", "AC-1 (2)", "AC-01 (02)", "ac-1.2", "ac-01.02"},
	{IDTypeNIST, "AC-01 (02)", "AC-1 (2)", "AC-01 (02)", "ac-1.2", "ac-01.02"},
	{IDTypeNIST, "PE-3", "PE-3", "PE-03", "pe-3", "pe-03"},
	{IDTypeNIST, "PE-03 (01)", "PE-3 (1)", "PE-03 (01)", "pe-3.1", "pe-03.01"},
}

func TestID(t *testing.T) {
	for _, tt := range idTests {
		var try ID
		var err error
		switch tt.vType {
		case IDTypeNIST:
			try, err = ParseIDFromNIST(tt.v)
			if err != nil {
				t.Errorf("rev5.ParseIDFromOSCAL(\"%s\") error: (%s)", tt.v, err.Error())
				continue
			}
		case IDTypeOSCAL:
			try, err = ParseIDFromOSCAL(tt.v)
			if err != nil {
				t.Errorf("rev5.ParseIDFromOSCAL(\"%s\") error: (%s)", tt.v, err.Error())
				continue
			}
		}
		if tryNIST, err := try.FormatNIST(); err != nil {
			t.Errorf("rev5.FormatNIST() error: on (%s) error (%s)", tt.v, err.Error())
		} else if tryNIST != tt.wantNIST {
			t.Errorf("rev5.FormatNIST() mimatch: on (%s) want (%s) got (%s)", tt.v, tt.wantNIST, tryNIST)
		}
		if tryNISTSort, err := try.FormatNISTSort(); err != nil {
			t.Errorf("rev5.FormatNISTSort() error: on (%s) error (%s)", tt.v, err.Error())
		} else if tryNISTSort != tt.wantNISTSort {
			t.Errorf("rev5.FormatNISTSort() mimatch: on (%s) want (%s) got (%s)", tt.v, tt.wantNISTSort, tryNISTSort)
		}
		if tryOSCAL, err := try.FormatOSCAL(); err != nil {
			t.Errorf("rev5.FormatOSCAL() error: on (%s) error (%s)", tt.v, err.Error())
		} else if tryOSCAL != tt.wantOSCAL {
			t.Errorf("rev5.FormatOSCAL() mimatch: on (%s) want (%s) got (%s)", tt.v, tt.wantOSCAL, tryOSCAL)
		}
		if tryOSCALSort, err := try.FormatOSCALSort(); err != nil {
			t.Errorf("rev5.FormatOSCALSort() error: on (%s) error (%s)", tt.v, err.Error())
		} else if tryOSCALSort != tt.wantOSCALSort {
			t.Errorf("rev5.FormatOSCALSort() mimatch: on (%s) want (%s) got (%s)", tt.v, tt.wantOSCALSort, tryOSCALSort)
		}
	}
}
