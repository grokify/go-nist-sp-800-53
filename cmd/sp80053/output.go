package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"slices"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

const (
	formatText = "text"
	formatJSON = "json"
	formatCSV  = "csv"
)

func addFormatFlag(cmd *cobra.Command, p *string, formats ...string) {
	cmd.Flags().StringVarP(p, "format", "f", formatText, "output format: "+joinFormats(formats))
}

func joinFormats(f []string) string {
	s := ""
	for i, v := range f {
		if i > 0 {
			s += ", "
		}
		s += v
	}
	return s
}

func checkFormat(format string, formats ...string) error {
	if !slices.Contains(formats, format) {
		return fmt.Errorf("unknown format %q: use %s", format, joinFormats(formats))
	}
	return nil
}

func writeJSON(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

// writeTable writes rows as aligned text or CSV.
func writeTable(w io.Writer, format string, header []string, rows [][]string) error {
	if format == formatCSV {
		cw := csv.NewWriter(w)
		if err := cw.Write(header); err != nil {
			return err
		}
		if err := cw.WriteAll(rows); err != nil {
			return err
		}
		return cw.Error()
	}
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	for _, r := range append([][]string{header}, rows...) {
		for i, c := range r {
			if i > 0 {
				if _, err := io.WriteString(tw, "\t"); err != nil {
					return err
				}
			}
			if _, err := io.WriteString(tw, c); err != nil {
				return err
			}
		}
		if _, err := io.WriteString(tw, "\n"); err != nil {
			return err
		}
	}
	return tw.Flush()
}
