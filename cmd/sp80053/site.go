package main

import (
	"fmt"

	"github.com/opensecurity-dev/go-nist-sp-800-53/rev5"
	"github.com/spf13/cobra"
)

func newSiteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "site",
		Short: "Maintain the MkDocs docs site",
	}
	cmd.AddCommand(newSiteGenerateCmd())
	return cmd
}

func newSiteGenerateCmd() *cobra.Command {
	var dir string
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Regenerate the control reference pages and mkdocs.yml",
		Long: `Write the per-baseline control reference pages to docs/rev5 and the site
navigation to mkdocs.yml, under --dir (the repository root).`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if err := (rev5.SiteWriter{}).WriteMarkdownPages(dir); err != nil {
				return err
			}
			_, err := fmt.Fprintf(cmd.OutOrStdout(), "wrote docs/rev5 and mkdocs.yml under %s\n", dir)
			return err
		},
	}
	cmd.Flags().StringVarP(&dir, "dir", "d", ".", "repository root")
	return cmd
}
