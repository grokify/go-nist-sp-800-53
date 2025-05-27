package rev5

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

type SiteWriter struct{}

func (sw SiteWriter) WriteMarkdownPagesTiers(baseDir string) error {
	tiers := []string{
		TierHighBaseline,
		TierHighUplift,
		TierModerateBaseline,
		TierModerateUplift,
		TierLowBaseline}

	for _, tierName := range tiers {
		err := sw.WriteMarkdownPagesTier(baseDir, tierName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sw SiteWriter) WriteMarkdownPagesTier(baseDir string, tierName string) error {
	tierDir := filepath.Join(baseDir, tierName)
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
		return fmt.Errorf("tierName not supported (%s)", tierName)
	}
	return sw.writeMarkdownPagesCatalog(tierDir, cat, idOSCALFilterIncl)
}

func (sw SiteWriter) writeMarkdownPagesCatalog(dir string, cat *Catalog, idOSCALFilterIncl []string) error {
	for _, group := range *cat.Groups {
		familyID := group.ID
		filename := familyID + ".md"
		fp := filepath.Join(dir, filename)
		if err := sw.writeMarkdownPagesCatalogGroup(fp, group, idOSCALFilterIncl); err != nil {
			return err
		}
	}
	return nil
}

func (sw SiteWriter) writeMarkdownPagesCatalogGroup(filename string, group oscalTypes.Group, idOSCALFilterIncl []string) error {
	if mdContent, err := sw.groupToStringsBuilder(group, idOSCALFilterIncl); err != nil {
		return err
	} else {
		return os.WriteFile(filename, []byte(mdContent.String()), 0600)
	}
}

func (sw SiteWriter) groupToStringsBuilder(group oscalTypes.Group, idOSCALFilterIncl []string) (*strings.Builder, error) {
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
			return nil, err
		} else if controlIDNIST, err := id.FormatNIST(); err != nil {
			return nil, err
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

	return &mdContent, nil
}
