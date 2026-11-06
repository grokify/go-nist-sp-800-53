package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/opensecurity-dev/nist-sp-800-53-go/rev5"
	"github.com/spf13/cobra"
)

var idTypes = map[string]rev5.IDType{
	"nist":       rev5.IDTypeNIST,
	"nist-sort":  rev5.IDTypeNISTSort,
	"oscal":      rev5.IDTypeOSCAL,
	"oscal-sort": rev5.IDTypeOSCALSort,
}

func newIDsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ids",
		Short: "Convert and check control IDs",
	}
	cmd.AddCommand(newIDsConvertCmd(), newIDsCheckCmd())
	return cmd
}

// readIDs returns args, or IDs read from stdin (one or more per line,
// separated by spaces or commas) when args is empty or "-".
func readIDs(in io.Reader, args []string) ([]string, error) {
	if len(args) > 0 && !(len(args) == 1 && args[0] == "-") {
		return args, nil
	}
	var out []string
	sc := bufio.NewScanner(in)
	for sc.Scan() {
		for _, f := range strings.FieldsFunc(sc.Text(), func(r rune) bool { return r == ',' || r == ' ' || r == '\t' }) {
			out = append(out, f)
		}
	}
	return out, sc.Err()
}

func newIDsConvertCmd() *cobra.Command {
	var to string
	cmd := &cobra.Command{
		Use:   "convert [id...]",
		Short: "Convert control IDs between NIST and OSCAL forms",
		Long: `Convert control IDs between forms:

  nist        AC-2 (1)
  nist-sort   AC-02 (01)
  oscal       ac-2.1
  oscal-sort  ac-02.01

With no arguments, or "-", IDs are read from stdin. NIST-form IDs that
contain a space, such as "AC-2 (1)", must be quoted.`,
		Example: "  sp80053 ids convert ac-2.1 AC-3 --to nist\n  cat ids.txt | sp80053 ids convert --to oscal-sort",
		RunE: func(cmd *cobra.Command, args []string) error {
			idType, ok := idTypes[to]
			if !ok {
				return fmt.Errorf("unknown --to %q: use nist, nist-sort, oscal or oscal-sort", to)
			}
			inputs, err := readIDs(cmd.InOrStdin(), args)
			if err != nil {
				return err
			}
			w := cmd.OutOrStdout()
			for _, s := range inputs {
				id, err := rev5.ParseID(s)
				if err != nil {
					return fmt.Errorf("%s: %w", s, err)
				}
				out, err := id.FormatType(idType)
				if err != nil {
					return fmt.Errorf("%s: %w", s, err)
				}
				if _, err := fmt.Fprintln(w, out); err != nil {
					return err
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&to, "to", "t", "nist", "target form: nist, nist-sort, oscal or oscal-sort")
	return cmd
}

func newIDsCheckCmd() *cobra.Command {
	var (
		format         string
		allowWithdrawn bool
	)
	cmd := &cobra.Command{
		Use:   "check [id...]",
		Short: "Check that control IDs exist in the catalog",
		Long: `Check that each control ID is well formed and names a control in the
catalog. It exits non-zero if any ID is invalid, unknown, or withdrawn
(unless --allow-withdrawn is set).

With no arguments, or "-", IDs are read from stdin.`,
		Example: "  sp80053 ids check AC-2 AU-2 SC-7\n  sp80053 ids check --format json < ids.txt",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := checkFormat(format, formatText, formatJSON); err != nil {
				return err
			}
			inputs, err := readIDs(cmd.InOrStdin(), args)
			if err != nil {
				return err
			}
			results := rev5.CatalogAll().CheckIDs(inputs)
			w := cmd.OutOrStdout()
			if format == formatJSON {
				err = writeJSON(w, results)
			} else {
				rows := make([][]string, 0, len(results))
				for _, r := range results {
					rows = append(rows, []string{r.Input, r.ID, r.Status, r.Error})
				}
				err = writeTable(w, formatText, []string{"INPUT", "ID", "STATUS", "DETAIL"}, rows)
			}
			if err != nil {
				return err
			}
			for _, r := range results {
				if r.Status == rev5.IDCheckOK || (allowWithdrawn && r.Status == rev5.IDCheckWithdrawn) {
					continue
				}
				return errFindings
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&allowWithdrawn, "allow-withdrawn", false, "accept withdrawn controls")
	addFormatFlag(cmd, &format, formatText, formatJSON)
	return cmd
}
