package rev5

import (
	"fmt"
	"strings"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

type IDType string

const (
	IDTypeOSCAL     IDType = "id"
	IDTypeOSCALSort IDType = "sort-id"
	IDTypeNIST      IDType = "label"
	IDTypeNISTSort  IDType = "sp800-53a"
)

type Control oscalTypes.Control

func (ctr Control) IDForType(idtype IDType) (string, error) {
	switch idtype {
	case IDTypeOSCAL:
		return ctr.ID, nil
	default:
		for _, prop := range *ctr.Props {
			if idtype == IDTypeOSCALSort {
				if prop.Name == string(idtype) {
					if v := strings.TrimSpace(prop.Value); v != "" {
						return v, nil
					}
				}
			} else if idtype == IDTypeNISTSort {
				if prop.Name == PropNameLabel && prop.Class == string(idtype) {
					if v := strings.TrimSpace(prop.Value); v != "" {
						return v, nil
					}
				}
			} else if idtype == IDTypeNIST {
				if prop.Name == PropNameLabel && prop.Class == "" {
					if v := strings.TrimSpace(prop.Value); v != "" {
						return v, nil
					}
				}
			} else {
				return "", fmt.Errorf("idtype not supported (%s)", idtype)
			}
		}
	}
	return "", fmt.Errorf("idtype not supported (%s)", idtype)
}
