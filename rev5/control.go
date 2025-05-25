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

	PropNameStatus  = "status"
	StatusWithdrawn = "withdrawn"
)

type Control oscalTypes.Control

func (ctr Control) Status() *string {
	if ctr.Props == nil {
		return nil
	}
	for _, prop := range *ctr.Props {
		if prop.Name == PropNameStatus {
			return &prop.Value
		}
	}
	return nil
}

// IDForType returns the control id for a requested status type.
func (ctr Control) IDForType(idtype IDType) (string, error) {
	if idtype == IDTypeOSCAL || strings.TrimSpace(string(idtype)) == "" {
		return ctr.ID, nil
	} else if ctr.Props == nil {
		return "", fmt.Errorf("no data found for found idtype (%s)", string(idtype))
	}
	for _, prop := range *ctr.Props {
		switch idtype {
		case IDTypeOSCALSort:
			if prop.Name == string(IDTypeOSCALSort) {
				if v := strings.TrimSpace(prop.Value); v != "" {
					return v, nil
				} else {
					return "", fmt.Errorf("no data found for found idtype (%s) for id (%s)", string(idtype), ctr.ID)
				}
			}
		case IDTypeNIST:
			fmt.Println("checking.NIST NO SORT")
			if prop.Name == PropNameLabel && strings.TrimSpace(prop.Class) == "" {
				if v := strings.TrimSpace(prop.Value); v != "" {
					return v, nil
				} else {
					return "", fmt.Errorf("no data found for found idtype (%s) for id (%s)", string(idtype), ctr.ID)
				}
			}
		case IDTypeNISTSort:
			if prop.Name == PropNameLabel && prop.Class == string(IDTypeNISTSort) {
				if v := strings.TrimSpace(prop.Value); v != "" {
					return v, nil
				} else {
					return "", fmt.Errorf("no data found for found idtype (%s) for id (%s)", string(idtype), ctr.ID)
				}
			}
		}
	}
	return "", fmt.Errorf("no key found for found idtype (%s)", string(idtype))
}
