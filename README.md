# NIST SP 800-53 Go SDK

[![Build Status][build-status-svg]][build-status-url]
[![Lint Status][lint-status-svg]][lint-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

Go SDK for NIST SP 800-53.

It provides high-level functions to access OSCAL data published by NIST. Notably, it:

1. provides higher-level functions to automate use cases
2. converts between various conrrol ID formats, e.g. NIST and OSCAL, with and without sort formatting
3. combines source data from NIST with the structs provided by Defense Unicorns
4. renders control documentation from OSCAL

## Credits

2. Data is sourced from [`github.com/usnistgov/oscal-content`](https://github.com/usnistgov/oscal-content).
1. DTOs are leveraged from [`github.com/defenseunicorns/go-oscal`](https://github.com/defenseunicorns/go-oscal).

 [build-status-svg]: https://github.com/grokify/go-nist-sp-800-53/actions/workflows/ci.yaml/badge.svg?branch=main
 [build-status-url]: https://github.com/grokify/go-nist-sp-800-53/actions/workflows/ci.yaml
 [lint-status-svg]: https://github.com/grokify/go-nist-sp-800-53/actions/workflows/lint.yaml/badge.svg?branch=main
 [lint-status-url]: https://github.com/grokify/go-nist-sp-800-53/actions/workflows/lint.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-nist-sp-800-53
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/go-nist-sp-800-53
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-nist-sp-800-53
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/go-nist-sp-800-53
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/go-nist-sp-800-53/blob/master/LICENSE
 [used-by-svg]: https://sourcegraph.com/github.com/grokify/go-nist-sp-800-53/-/badge.svg
 [used-by-url]: https://sourcegraph.com/github.com/grokify/go-nist-sp-800-53?badge
 [loc-svg]: https://tokei.rs/b1/github/grokify/go-nist-sp-800-53
 [repo-url]: https://github.com/grokify/go-nist-sp-800-53