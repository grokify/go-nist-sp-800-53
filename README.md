# NIST SP 800-53 Go SDK

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Coverage][coverage-svg]][coverage-url]
[![License][license-svg]][license-url]

Go SDK and CLI for NIST SP 800-53 Rev. 5, built on the OSCAL catalog and baselines published by NIST.

It:

1. embeds the NIST SP 800-53 Rev. 5 OSCAL catalog and the low, moderate and high baseline profiles
2. looks up controls by baseline, family and implementation level (organization or system)
3. converts control IDs between NIST (`AC-2 (1)`) and OSCAL (`ac-2.1`) forms, with and without sort padding
4. checks that control IDs name real, non-withdrawn controls
5. renders a control reference site, published at [opensecurity-dev.github.io/go-nist-sp-800-53](https://opensecurity-dev.github.io/go-nist-sp-800-53)

## Install

```bash
go get github.com/opensecurity-dev/go-nist-sp-800-53
go install github.com/opensecurity-dev/go-nist-sp-800-53/cmd/sp80053@latest
```

## CLI

```bash
sp80053 baselines                                    # controls per baseline
sp80053 families                                     # families with per-baseline counts
sp80053 control show 'AC-2(1)'                       # statement, guidance, baselines, level
sp80053 controls list --baseline moderate --level system --format csv
sp80053 controls list --family ac,au --ids-only
sp80053 ids convert ac-2.1 AC-3 --to oscal-sort      # ac-02.01, ac-03
sp80053 ids check AC-2 AU-99                         # exits non-zero on unknown IDs
sp80053 site generate                                # regenerate docs/rev5 and mkdocs.yml
```

Most commands take `--format text|json|csv`. `ids convert` and `ids check` read IDs from stdin when none are given.
`--level system` lists controls a system implements itself; `--level organization` lists organization-wide controls such
as policies and procedures.

## Library

```go
import "github.com/opensecurity-dev/go-nist-sp-800-53/rev5"

cat := rev5.CatalogAll()

// Moderate-baseline access-control controls implemented at the system level.
controls, err := cat.Query(rev5.Query{
    Baseline: rev5.BaselineModerate,
    Families: []string{"AC"},
    Level:    rev5.PropImplementationLevelValueSystem,
})

// One control, with its statement and guidance.
info, err := cat.ControlInfo("AC-2(1)")

// Check IDs, for example the control references in a document.
results := cat.CheckIDs([]string{"AC-2", "AU-99"})
```

## Docs site

The MkDocs source is in `docs/`. The control pages under `docs/rev5` and the navigation in `mkdocs.yml` are generated;
CI builds the site and publishes it to GitHub Pages.

```bash
go run ./cmd/sp80053 site generate   # regenerate the control pages
pip install -r docs/requirements.txt
mkdocs serve                          # preview locally
```

## Credits

1. Data is sourced from [`github.com/usnistgov/oscal-content`](https://github.com/usnistgov/oscal-content).
2. OSCAL types come from [`github.com/defenseunicorns/go-oscal`](https://github.com/defenseunicorns/go-oscal).

 [go-ci-svg]: https://github.com/opensecurity-dev/go-nist-sp-800-53/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/opensecurity-dev/go-nist-sp-800-53/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/opensecurity-dev/go-nist-sp-800-53/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/opensecurity-dev/go-nist-sp-800-53/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/opensecurity-dev/go-nist-sp-800-53/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/opensecurity-dev/go-nist-sp-800-53/actions/workflows/go-sast-codeql.yaml
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/opensecurity-dev/go-nist-sp-800-53
 [docs-godoc-url]: https://pkg.go.dev/github.com/opensecurity-dev/go-nist-sp-800-53
 [coverage-svg]: https://img.shields.io/badge/coverage-43.8%25-red
 [coverage-url]: https://github.com/opensecurity-dev/go-nist-sp-800-53
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/opensecurity-dev/go-nist-sp-800-53/blob/main/LICENSE
