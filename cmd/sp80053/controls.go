package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/opensecurity-dev/nist-sp-800-53-go/rev5"
	"github.com/spf13/cobra"
)

func newControlsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "controls",
		Short: "List and filter controls",
	}
	cmd.AddCommand(newControlsListCmd())
	return cmd
}

func newControlsListCmd() *cobra.Command {
	var (
		q       rev5.Query
		format  string
		idsOnly bool
	)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List controls by baseline, family and implementation level",
		Long: `List controls and enhancements, in catalog order. Withdrawn controls are
left out unless --include-withdrawn is set.

--level system lists controls a system implements itself; --level
organization lists organization-wide controls such as policies.`,
		Example: "  sp80053 controls list --baseline moderate --level system\n" +
			"  sp80053 controls list --family ac,au --format csv\n" +
			"  sp80053 controls list --baseline low --ids-only",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := checkFormat(format, formatText, formatJSON, formatCSV); err != nil {
				return err
			}
			infos, err := rev5.CatalogAll().Query(q)
			if err != nil {
				return err
			}
			w := cmd.OutOrStdout()
			if idsOnly {
				for _, i := range infos {
					if _, err := fmt.Fprintln(w, i.ID); err != nil {
						return err
					}
				}
				return nil
			}
			if format == formatJSON {
				return writeJSON(w, infos)
			}
			header := []string{"ID", "TITLE", "BASELINES", "LEVEL", "WITHDRAWN"}
			rows := make([][]string, 0, len(infos))
			for _, i := range infos {
				rows = append(rows, []string{i.ID, i.Title, strings.Join(i.Baselines, ","),
					strings.Join(i.ImplementationLevels, ","), strconv.FormatBool(i.Withdrawn)})
			}
			return writeTable(w, format, header, rows)
		},
	}
	f := cmd.Flags()
	f.StringVarP(&q.Baseline, "baseline", "b", "", "baseline: low, moderate or high")
	f.StringSliceVar(&q.Families, "family", nil, "family abbreviations, such as ac,au")
	f.StringVarP(&q.Level, "level", "l", "", "implementation level: organization or system")
	f.BoolVar(&q.IncludeWithdrawn, "include-withdrawn", false, "include withdrawn controls")
	f.BoolVar(&idsOnly, "ids-only", false, "print only the control IDs, one per line")
	addFormatFlag(cmd, &format, formatText, formatJSON, formatCSV)
	return cmd
}
