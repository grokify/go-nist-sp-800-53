package main

import (
	"fmt"
	"strings"

	"github.com/opensecurity-dev/go-nist-sp-800-53/rev5"
	"github.com/spf13/cobra"
)

func newControlCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "control",
		Short: "Show a single control",
	}
	cmd.AddCommand(newControlShowCmd())
	return cmd
}

func newControlShowCmd() *cobra.Command {
	var format string
	cmd := &cobra.Command{
		Use:     "show <id>",
		Short:   "Show a control's baselines, implementation level, statement and guidance",
		Example: "  sp80053 control show AC-2(1)\n  sp80053 control show ac-2.1 --format json",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := checkFormat(format, formatText, formatJSON); err != nil {
				return err
			}
			info, err := rev5.CatalogAll().ControlInfo(args[0])
			if err != nil {
				if info.ID == "" {
					return err
				}
				// The control exists but its statement could not be fully
				// rendered; show the rest and say so.
				fmt.Fprintln(cmd.ErrOrStderr(), "warning:", err)
			}
			w := cmd.OutOrStdout()
			if format == formatJSON {
				return writeJSON(w, info)
			}
			return writeControlText(w, info)
		},
	}
	addFormatFlag(cmd, &format, formatText, formatJSON)
	return cmd
}

func writeControlText(w interface{ Write([]byte) (int, error) }, info rev5.ControlInfo) error {
	var b strings.Builder
	fmt.Fprintf(&b, "%s: %s\n\n", info.ID, info.Title)
	fmt.Fprintf(&b, "OSCAL ID:              %s\n", info.OSCALID)
	fmt.Fprintf(&b, "Family:                %s\n", info.Family)
	if info.Withdrawn {
		b.WriteString("Status:                withdrawn\n")
	}
	fmt.Fprintf(&b, "Baselines:             %s\n", orNone(info.Baselines))
	fmt.Fprintf(&b, "Implementation level:  %s\n", orNone(info.ImplementationLevels))
	if info.Statement != "" {
		fmt.Fprintf(&b, "\nStatement:\n\n%s\n", indent(info.Statement))
	}
	if info.Guidance != "" {
		fmt.Fprintf(&b, "\nGuidance:\n\n%s\n", indent(info.Guidance))
	}
	_, err := w.Write([]byte(b.String()))
	return err
}

func indent(s string) string {
	return "  " + strings.ReplaceAll(s, "\n", "\n  ")
}

func orNone(s []string) string {
	if len(s) == 0 {
		return "none"
	}
	return strings.Join(s, ", ")
}
