package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

func run(t *testing.T, stdin string, args ...string) (string, error) {
	t.Helper()
	cmd := newRootCmd()
	var out, errOut bytes.Buffer
	cmd.SetOut(&out)
	cmd.SetErr(&errOut)
	cmd.SetIn(strings.NewReader(stdin))
	cmd.SetArgs(args)
	err := cmd.Execute()
	return out.String(), err
}

func TestBaselines(t *testing.T) {
	out, err := run(t, "", "baselines", "--format", "csv")
	if err != nil {
		t.Fatal(err)
	}
	want := "baseline,controls,added\nlow,149,149\nmoderate,287,138\nhigh,370,83\n"
	if !strings.EqualFold(out, want) {
		t.Errorf("got:\n%s\nwant:\n%s", out, want)
	}
}

func TestFamiliesJSON(t *testing.T) {
	out, err := run(t, "", "families", "-f", "json")
	if err != nil {
		t.Fatal(err)
	}
	var fams []struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal([]byte(out), &fams); err != nil {
		t.Fatal(err)
	}
	if len(fams) != 20 || fams[0].ID != "AC" {
		t.Errorf("unexpected families: %v", fams)
	}
}

func TestControlShow(t *testing.T) {
	out, err := run(t, "", "control", "show", "AC-2(1)")
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{"AC-2 (1): Automated System Account Management", "Baselines:             moderate, high", "Statement:", "Guidance:"} {
		if !strings.Contains(out, s) {
			t.Errorf("output missing %q:\n%s", s, out)
		}
	}
	if _, err := run(t, "", "control", "show", "AC-99"); err == nil {
		t.Error("show AC-99 succeeded, want an error")
	}
}

func TestControlsList(t *testing.T) {
	out, err := run(t, "", "controls", "list", "-b", "low", "--family", "ac", "-l", "system", "--ids-only")
	if err != nil {
		t.Fatal(err)
	}
	ids := strings.Fields(out)
	if len(ids) == 0 || ids[0] != "AC-3" {
		t.Errorf("ids = %v, want AC-3 first", ids)
	}
	if _, err := run(t, "", "controls", "list", "-b", "extreme"); err == nil {
		t.Error("unknown baseline accepted")
	}
	if _, err := run(t, "", "controls", "list", "-f", "xml"); err == nil {
		t.Error("unknown format accepted")
	}
}

func TestIDsConvert(t *testing.T) {
	out, err := run(t, "", "ids", "convert", "ac-2.1", "AC-3", "--to", "oscal-sort")
	if err != nil {
		t.Fatal(err)
	}
	if out != "ac-02.01\nac-03\n" {
		t.Errorf("got %q", out)
	}
	out, err = run(t, "ac-2.1, sc-7\n", "ids", "convert")
	if err != nil {
		t.Fatal(err)
	}
	if out != "AC-2 (1)\nSC-7\n" {
		t.Errorf("stdin: got %q", out)
	}
	if _, err := run(t, "", "ids", "convert", "XX-1"); err == nil {
		t.Error("invalid ID accepted")
	}
}

func TestIDsCheck(t *testing.T) {
	if _, err := run(t, "", "ids", "check", "AC-2", "SC-7"); err != nil {
		t.Errorf("valid IDs: %v", err)
	}
	out, err := run(t, "", "ids", "check", "AC-2", "AC-99")
	if !errors.Is(err, errFindings) {
		t.Errorf("unknown ID: err = %v, want errFindings", err)
	}
	if !strings.Contains(out, "unknown") {
		t.Errorf("output should report unknown:\n%s", out)
	}
	if _, err := run(t, "", "ids", "check", "AC-13"); !errors.Is(err, errFindings) {
		t.Error("withdrawn ID should fail by default")
	}
	if _, err := run(t, "", "ids", "check", "AC-13", "--allow-withdrawn"); err != nil {
		t.Errorf("--allow-withdrawn: %v", err)
	}
}
