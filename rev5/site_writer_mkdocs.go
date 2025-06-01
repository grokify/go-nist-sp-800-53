package rev5

import mkdocs "github.com/grokify/go-mkdocs"

func (sw SiteWriter) mkdocsConfigDefault() mkdocs.Config {
	return mkdocs.Config{
		SiteName: "NIST SP 800-53 Control Reference Site",
		SiteURL:  "https://grokify.github.io/go-nist-sp-800-53",
		DocsDir:  "docsrc",
		Theme:    "readthedocs",
		Nav: mkdocs.Texts{
			{
				Key: "Home",
				Val: "index.md",
			},
		}}
}

func (sw SiteWriter) writeMkdocsTOC(filename string, navTexts mkdocs.Texts) error {
	cfg := sw.mkdocsConfigDefault()
	cfg.Nav = append(cfg.Nav, navTexts...)
	return cfg.WriteFileYAML(filename, 0600)
}
