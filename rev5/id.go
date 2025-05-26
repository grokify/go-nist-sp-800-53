package rev5

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type ID struct {
	OSCALSortID        string
	FamilyAbbreviation string
	BaseControlNumber  int
	EnhancementNumber  int
}

var rxControlBaseOnly = regexp.MustCompile(`^([a-zA-Z]{2})\-([0-9]{1,2})\s*(\(([0-9]{1,2})\))?$`)

func ParseID(s string) (ID, error) {
	if strings.Contains(s, ".") {
		return parseIDFromOSCAL(s)
	} else {
		return parseIDFromNIST(s)
	}
}

func parseIDFromNIST(s string) (ID, error) {
	id := ID{}
	s = strings.ToUpper(strings.TrimSpace(s))
	if s == "" {
		return id, errors.New("cannot parse empty string")
	}
	m := rxControlBaseOnly.FindStringSubmatch(s)
	if len(m) == 0 {
		return id, fmt.Errorf("cannot parse id format for input (%s)", s)
	}
	if cfAbbr := strings.ToLower(m[1]); !IsFamilyAbbreviation(cfAbbr) {
		return id, fmt.Errorf("id family is unknown (%s)", cfAbbr)
	} else {
		id.FamilyAbbreviation = cfAbbr
	}
	if intval, err := strconv.Atoi(m[2]); err != nil {
		return id, err
	} else {
		id.BaseControlNumber = intval
	}
	if m[4] != "" {
		if intval, err := strconv.Atoi(m[4]); err != nil {
			return id, err
		} else {
			id.EnhancementNumber = intval
		}
	}
	if err := id.idPartsVerify(); err != nil {
		return ID{}, err
	} else if oscalSortID, err := id.FormatOSCALSort(); err != nil {
		return ID{}, err
	} else {
		id.OSCALSortID = oscalSortID
		return id, nil
	}
}

func parseIDFromOSCAL(s string) (ID, error) {
	id := ID{}
	s = strings.TrimSpace(s)
	if s == "" {
		return id, errors.New("cannot parse empty string")
	}
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return id, fmt.Errorf("id has wrong format with (%d) hypen parts", len(parts))
	}
	if cfAbbr := strings.ToLower(parts[0]); !IsFamilyAbbreviation(cfAbbr) {
		return id, fmt.Errorf("id family is unknown (%s)", cfAbbr)
	} else {
		id.FamilyAbbreviation = cfAbbr
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
	if err := id.idPartsVerify(); err != nil {
		return ID{}, err
	} else if oscalSortID, err := id.FormatOSCALSort(); err != nil {
		return ID{}, err
	} else {
		id.OSCALSortID = oscalSortID
		return id, nil
	}
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

func (id ID) idPartsVerify() error {
	if _, _, _, err := id.idParts(); err != nil { //nolint:dogsled
		return err
	} else {
		return nil
	}
}

func (id ID) idParts() (string, int, int, error) {
	cfa := strings.ToLower(strings.TrimSpace(id.FamilyAbbreviation))
	if !IsFamilyAbbreviation(cfa) {
		return "", -1, -1, fmt.Errorf("control family abbreviation (%s) not set or is unknown", cfa)
	}
	if id.BaseControlNumber < 1 {
		return "", -1, -1, errors.New("base control number is not value")
	}
	return cfa, id.BaseControlNumber, id.EnhancementNumber, nil
}

func (id ID) formatInternal(layoutBase, layoutEnhancement string) (string, error) {
	if cfa, base, enh, err := id.idParts(); err != nil {
		return "", err
	} else if enh < 1 {
		return fmt.Sprintf(layoutBase, cfa, base), nil
	} else {
		return fmt.Sprintf(layoutEnhancement, cfa, base, enh), nil
	}
}

func (id ID) FormatNIST() (string, error) {
	if s, err := id.formatInternal(layoutControlIDNISTBase, layoutControlIDNISTEnh); err != nil {
		return "", err
	} else {
		return strings.ToUpper(s), nil
	}
}

func (id ID) FormatNISTSort() (string, error) {
	if s, err := id.formatInternal(layoutControlIDNISTBaseSort, layoutControlIDNISTEnhSort); err != nil {
		return "", err
	} else {
		return strings.ToUpper(s), nil
	}
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
