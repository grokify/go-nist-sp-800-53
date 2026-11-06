package main

import (
	"strconv"

	"github.com/opensecurity-dev/nist-sp-800-53-go/rev5"
	"github.com/spf13/cobra"
)

func newFamiliesCmd() *cobra.Command {
	var format string
	cmd := &cobra.Command{
		Use:   "families",
		Short: "List control families with control counts per baseline",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := checkFormat(format, formatText, formatJSON, formatCSV); err != nil {
				return err
			}
			fams, err := rev5.CatalogAll().FamilyInfos()
			if err != nil {
				return err
			}
			w := cmd.OutOrStdout()
			if format == formatJSON {
				return writeJSON(w, fams)
			}
			header := []string{"ID", "TITLE", "CONTROLS", "LOW", "MODERATE", "HIGH"}
			rows := make([][]string, 0, len(fams))
			for _, f := range fams {
				rows = append(rows, []string{f.ID, f.Title, strconv.Itoa(f.Controls),
					strconv.Itoa(f.Baselines[rev5.BaselineLow]),
					strconv.Itoa(f.Baselines[rev5.BaselineModerate]),
					strconv.Itoa(f.Baselines[rev5.BaselineHigh])})
			}
			return writeTable(w, format, header, rows)
		},
	}
	addFormatFlag(cmd, &format, formatText, formatJSON, formatCSV)
	return cmd
}
