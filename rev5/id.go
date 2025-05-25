package rev5

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ID struct {
	ControlFamilyAbbrivation string
	BaseControlNumber        int
	EnhancementNumber        int
}

func ParseIDFromOSCAL(s string) (ID, error) {
	id := ID{}
	s = strings.TrimSpace(s)
	if s == "" {
		return id, errors.New("cannot parse empty string")
	}
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return id, fmt.Errorf("id has wrong format with (%d) hypen parts", len(parts))
	}
	cfAbbr := strings.ToUpper(parts[0])
	if !IsControlFamilyAbbreviation(cfAbbr) {
		return id, fmt.Errorf("id family is unknown (%s)", cfAbbr)
	} else {
		id.ControlFamilyAbbrivation = cfAbbr
	}
	numParts := strings.Split(parts[1], ".")
	if len(numParts) < 1 || len(numParts) > 2 {
		return id, fmt.Errorf("id number has wrong format with (%d) period parts", len(numParts))
	}
	baseNum, err := strconv.Atoi(numParts[0])
	if err != nil {
		return id, err
	} else {
		id.BaseControlNumber = baseNum
	}
	if len(numParts) > 1 {
		enhNum, err := strconv.Atoi(numParts[1])
		if err != nil {
			return id, err
		} else {
			id.EnhancementNumber = enhNum
		}
	}
	return id, nil
}

const (
	layoutControlOSCALBase      = "%s-%d"
	layoutControlOSCALEnh       = "%s-%d.%d"
	layoutControlOSCALBaseSort  = "%s-%02d"
	layoutControlOSCALEnhSort   = "%s-%02d.%02d"
	layoutControlIDNISTBase     = "%s-%d"
	layoutControlIDNISTEnh      = "%s-%d (%d)"
	layoutControlIDNISTBaseSort = "%s-%02d"
	layoutControlIDNISTEnhSort  = "%s-%02d (%02d)"
)

func (id ID) idPartsVerified() (string, int, int, error) {
	cfa := strings.ToUpper(strings.TrimSpace(id.ControlFamilyAbbrivation))
	if !IsControlFamilyAbbreviation(cfa) {
		return "", -1, -1, fmt.Errorf("control family abbreviation (%s) not set or is unknown", cfa)
	}
	if id.BaseControlNumber < 1 {
		return "", -1, -1, errors.New("base control number is not value")
	}
	return cfa, id.BaseControlNumber, id.EnhancementNumber, nil
}

func (id ID) formatInternal(layoutBase, layoutEnhancement string) (string, error) {
	if cfa, base, enh, err := id.idPartsVerified(); err != nil {
		return "", err
	} else if enh < 1 {
		return fmt.Sprintf(layoutBase, cfa, base), nil
	} else {
		return fmt.Sprintf(layoutEnhancement, cfa, base, enh), nil
	}
}

func (id ID) FormatNIST() (string, error) {
	return id.formatInternal(layoutControlIDNISTBase, layoutControlIDNISTEnh)
}

func (id ID) FormatNISTSort() (string, error) {
	return id.formatInternal(layoutControlIDNISTBaseSort, layoutControlIDNISTEnhSort)
}

func (id ID) FormatOSCAL() (string, error) {
	if s, err := id.formatInternal(layoutControlOSCALBase, layoutControlOSCALEnh); err != nil {
		return "", err
	} else {
		return strings.ToLower(s), nil
	}
}

func (id ID) FormatOSCALSort() (string, error) {
	if s, err := id.formatInternal(layoutControlOSCALBaseSort, layoutControlOSCALEnhSort); err != nil {
		return "", err
	} else {
		return strings.ToLower(s), nil
	}
}
