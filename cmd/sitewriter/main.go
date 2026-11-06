// Command sitewriter regenerates the MkDocs site source: the control
// reference pages under docs/rev5 and the navigation in mkdocs.yml.
//
// Run it from the repository root:
//
//	go run ./cmd/sitewriter
package main

import (
	"flag"
	"log"

	"github.com/opensecurity-dev/go-nist-sp-800-53/rev5"
)

func main() {
	dir := flag.String("dir", ".", "repository root to write docs/ and mkdocs.yml into")
	flag.Parse()

	if err := (rev5.SiteWriter{}).WriteMarkdownPages(*dir); err != nil {
		log.Fatal(err)
	}
}
