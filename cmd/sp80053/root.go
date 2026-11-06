package main

import (
	"errors"

	"github.com/spf13/cobra"
)

// errFindings means a command ran and reported problems itself, such as
// unknown control IDs. main exits non-zero without printing it again.
var errFindings = errors.New("findings reported")

func newRootCmd() *cobra.Command {
	root := &cobra.Command{
		Use:   "sp80053",
		Short: "Look up NIST SP 800-53 Rev. 5 controls",
		Long: `sp80053 looks up NIST SP 800-53 Rev. 5 controls, baselines and families
from the NIST OSCAL catalog embedded in the binary, converts and checks
control IDs, and regenerates the docs site.

Control IDs may be written in NIST form (AC-2(1)) or OSCAL form (ac-2.1).`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.AddCommand(
		newFamiliesCmd(),
		newControlCmd(),
		newControlsCmd(),
		newIDsCmd(),
		newBaselinesCmd(),
		newSiteCmd(),
	)
	return root
}
