package main

import (
	"strconv"

	"github.com/opensecurity-dev/go-nist-sp-800-53/rev5"
	"github.com/spf13/cobra"
)

type baselineCount struct {
	Baseline string `json:"baseline"`
	Controls int    `json:"controls"`
	// Added counts controls not in the next lower baseline.
	Added int `json:"added"`
}

func newBaselinesCmd() *cobra.Command {
	var format string
	cmd := &cobra.Command{
		Use:   "baselines",
		Short: "Count the controls in each baseline",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := checkFormat(format, formatText, formatJSON, formatCSV); err != nil {
				return err
			}
			ids := rev5.ControlIDs{}
			counts := []baselineCount{
				{rev5.BaselineLow, len(ids.LowBaseline()), len(ids.LowBaseline())},
				{rev5.BaselineModerate, len(ids.ModerateBaseline()), len(ids.ModerateUplift())},
				{rev5.BaselineHigh, len(ids.HighBaseline()), len(ids.HighUplift())},
			}
			w := cmd.OutOrStdout()
			if format == formatJSON {
				return writeJSON(w, counts)
			}
			rows := make([][]string, 0, len(counts))
			for _, c := range counts {
				rows = append(rows, []string{c.Baseline, strconv.Itoa(c.Controls), strconv.Itoa(c.Added)})
			}
			return writeTable(w, format, []string{"BASELINE", "CONTROLS", "ADDED"}, rows)
		},
	}
	addFormatFlag(cmd, &format, formatText, formatJSON, formatCSV)
	return cmd
}
