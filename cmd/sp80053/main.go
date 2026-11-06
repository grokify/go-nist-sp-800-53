// Command sp80053 looks up NIST SP 800-53 Rev. 5 controls, baselines and
// families from the embedded OSCAL catalog, converts and checks control IDs,
// and regenerates the docs site.
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if err := newRootCmd().Execute(); err != nil {
		if !errors.Is(err, errFindings) {
			fmt.Fprintln(os.Stderr, "error:", err)
		}
		os.Exit(1)
	}
}
