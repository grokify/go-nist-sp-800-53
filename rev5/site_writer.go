package rev5

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
	mkdocs "github.com/grokify/go-mkdocs"
)

type SiteWriter struct{}

func (sw SiteWriter) WriteMarkdownPagesTiers(baseDir string) error {
	tiers := []string{
		TierHighBaseline,
		TierHighUplift,
		TierModerateBaseline,
		TierModerateUplift,
		TierLowBaseline}

	mkdocsTexts := mkdocs.Texts{}

	for _, tierName := range tiers {
		if tierMkdocsTexts, err := sw.WriteMarkdownPagesTier(baseDir, tierName); err != nil {
			return err
		} else {
			mkdocsTexts = append(mkdocsTexts, mkdocs.Text{
				Key:      TierDisplayOrDefault(tierName, tierName),
				Children: tierMkdocsTexts,
			})
		}
	}
	return sw.mkdocsWriteTOC(filepath.Join(baseDir, mkdocs.FilenameTOC), mkdocsTexts)
}

func (sw SiteWriter) WriteMarkdownPagesTier(baseDir string, tierName string) (mkdocs.Texts, error) {
	dirTierWrite := filepath.Join(baseDir, "docsrc/rev5", tierName)
	dirTierMkdocs := filepath.Join("rev5", tierName)
	var cat *Catalog
	var idOSCALFilterIncl []string
	ids := NewControlIDs()
	switch tierName {
	case TierHighBaseline:
		cat = CatalogHighBaseline()
	case TierHighUplift:
		cat = CatalogHighBaseline()
		idOSCALFilterIncl = ids.HighUplift()
	case TierModerateBaseline:
		cat = CatalogModerateBaseline()
	case TierModerateUplift:
		cat = CatalogModerateBaseline()
		idOSCALFilterIncl = ids.ModerateUplift()
	case TierLowBaseline:
		cat = CatalogLowBaseline()
	default:
		return mkdocs.Texts{}, fmt.Errorf("tierName not supported (%s)", tierName)
	}
	mkTexts, err := sw.writeMarkdownPagesCatalog(dirTierWrite, dirTierMkdocs, cat, idOSCALFilterIncl)
	if err != nil {
		return mkTexts, err
	}
	tierIndexFilename := filepath.Join(dirTierWrite, "index.md")
	return mkTexts, sw.writeMarkdownPageTierIndex(tierIndexFilename, tierName, mkTexts)
}

func (sw SiteWriter) writeMarkdownPageTierIndex(filename string, tierName string, mkTexts mkdocs.Texts) error {
	var mdContent strings.Builder
	tierDisplayname := TierDisplayOrDefault(tierName, tierName)
	ids := NewControlIDs()
	tierControls, err := ids.Tier(tierName)
	if err != nil {
		return err
	}
	mdContent.WriteString(fmt.Sprintf("# %s (%d)\n\n", tierDisplayname, len(tierControls)))
	for _, mktext := range mkTexts {
		u := mktext.Val
		_, filenameonly := filepath.Split(u)
		mdContent.WriteString(fmt.Sprintf("1. [%s](%s)\n", mktext.Key, filenameonly))
	}
	return os.WriteFile(filename, []byte(mdContent.String()), 0600)
}

func (sw SiteWriter) writeMarkdownPagesCatalog(dirWrite, dirMkdocs string, cat *Catalog, idOSCALFilterIncl []string) (mkdocs.Texts, error) {
	mkdocsTexts := mkdocs.Texts{}
	for _, group := range *cat.Groups {
		familyID := group.ID
		filename := familyID + ".md"
		fp := filepath.Join(dirWrite, filename)
		if toc, err := sw.writeMarkdownPagesCatalogGroup(fp, group, idOSCALFilterIncl); err != nil {
			return mkdocsTexts, err
		} else {
			mkdocsTexts = append(mkdocsTexts, mkdocs.Text{
				Key: toc,
				Val: filepath.Join(dirMkdocs, filename),
			})
		}
	}
	return mkdocsTexts, nil
}

func (sw SiteWriter) writeMarkdownPagesCatalogGroup(filename string, group oscalTypes.Group, idOSCALFilterIncl []string) (string, error) {
	if mdContent, familyTOCTitle, err := sw.groupToStringsBuilder(group, idOSCALFilterIncl); err != nil {
		return "", err
	} else {
		return familyTOCTitle, os.WriteFile(filename, []byte(mdContent.String()), 0600)
	}
}

func (sw SiteWriter) groupToStringsBuilder(group oscalTypes.Group, idOSCALFilterIncl []string) (*strings.Builder, string, error) {
	familyID := group.ID
	familyTitle := group.Title

	/*
		var mdBaseControls []string
		var mdEnhancements []string
		// Iterate over controls in the group
		for _, control := range *group.Controls {
			controlID := control.ID
			controlTitle := control.Title
			parts := Parts(*control.Parts)
			controlDesc := parts.ExtractProseString()

			// Check if the control is a base control or an enhancement
			//if strings.Contains(controlID, ".") {
			if control.Class == ClassSP80053Enhancement {
				// Enhancement
				mdEnhancements = append(mdEnhancements, fmt.Sprintf("#### %s: %s\n\n%s\n", controlID, controlTitle, controlDesc))
			} else {
				// Base control
				mdBaseControls = append(mdBaseControls, fmt.Sprintf("### %s: %s\n\n%s\n", controlID, controlTitle, controlDesc))
			}
		}
	*/

	var mdAllControls []string
	controls2 := Controls(*group.Controls)
	controlsFlattened := controls2.Flatten()
	if len(idOSCALFilterIncl) > 0 {
		controlsFlattened = controlsFlattened.Filter(idOSCALFilterIncl)
	}

	for _, control := range controlsFlattened {
		controlID := control.ID
		if id, err := ParseID(controlID); err != nil {
			return nil, "", err
		} else if controlIDNIST, err := id.FormatNIST(); err != nil {
			return nil, "", err
		} else {
			controlID = controlIDNIST
		}
		controlTitle := control.Title
		parts := Parts(*control.Parts)
		controlDesc := parts.ExtractProseString()
		mdAllControls = append(mdAllControls, fmt.Sprintf("### %s: %s\n\n%s\n", controlID, controlTitle, controlDesc))
	}

	// Create Markdown content
	var mdContent strings.Builder
	mdContent.WriteString(fmt.Sprintf("# %s - %s\n\n", strings.ToUpper(familyID), familyTitle))
	mdContent.WriteString(fmt.Sprintf("* Controls: %d\n\n", len(mdAllControls)))

	/*
			mdContent.WriteString(fmt.Sprintf("* Base Controls: %d\n\n", len(mdBaseControls)))
			mdContent.WriteString(fmt.Sprintf("* Control Enhancements: %d\n\n", len(mdEnhancements)))
			mdContent.WriteString("## Base Controls\n\n")
			for _, bc := range mdBaseControls {
				mdContent.WriteString(bc + "\n")
			}
			mdContent.WriteString("## Control Enhancements\n\n")
			for _, enh := range mdEnhancements {
				mdContent.WriteString(enh + "\n")
			}
		}
	*/
	mdContent.WriteString("## Controls\n\n")
	for _, mdControl := range mdAllControls {
		mdContent.WriteString(mdControl + "\n")
	}

	tocTitle := fmt.Sprintf("%s: %s (%d)", strings.ToUpper(familyID), familyTitle, len(controlsFlattened))
	return &mdContent, tocTitle, nil
}
