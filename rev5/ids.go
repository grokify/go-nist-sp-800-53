package rev5

import (
	"errors"
	"slices"
	"sort"
)

var ErrIDTypeNotSupported = errors.New("idtype not supported")

type IDs []ID

func ParseIDs(s []string) (IDs, error) {
	ids := IDs{}
	seen := map[string]int{}
	for _, si := range s {
		id, err := ParseID(si)
		if err != nil {
			return ids, err
		} else if oscalID, err := id.FormatOSCALSort(); err != nil {
			return ids, err
		} else if _, ok := seen[oscalID]; ok {
			continue
		} else {
			ids = append(ids, id)
			seen[oscalID]++
		}
	}
	ids.Sort()
	return ids, nil
}

func (ids IDs) FilterByTier(tierName string, ctrIDs *ControlIDs) (IDs, error) {
	if ctrIDs == nil {
		ctrIDs = NewControlIDs()
	}
	inclIDsRaw, err := ctrIDs.Tier(tierName)
	if err != nil {
		return IDs{}, err
	} else if len(inclIDsRaw) == 0 {
		return IDs{}, nil
	}
	inclIDs, err := ParseIDs(inclIDsRaw)
	if err != nil {
		return IDs{}, err
	}
	inclMap, err := inclIDs.formatOSCALSortMap()
	if err != nil {
		return IDs{}, err
	}
	var out IDs
	for _, id := range ids {
		if _, ok := inclMap[id.OSCALSortID]; ok {
			out = append(out, id)
		}
	}
	return out, nil
}

func (ids IDs) Sort() {
	slices.SortFunc(ids, func(i, j ID) int {
		if i.OSCALSortID < j.OSCALSortID {
			return -1
		} else if i.OSCALSortID > j.OSCALSortID {
			return 1
		} else {
			return 0
		}
	})
}

func (ids IDs) FormatNIST() ([]string, error) {
	return ids.formatFunc(func(id ID) (string, error) { return id.FormatNIST() })
}

func (ids IDs) FormatNISTSort() ([]string, error) {
	return ids.formatFunc(func(id ID) (string, error) { return id.FormatNISTSort() })
}

func (ids IDs) FormatOSCAL() ([]string, error) {
	return ids.formatFunc(func(id ID) (string, error) { return id.FormatOSCAL() })
}

func (ids IDs) FormatOSCALSort() ([]string, error) {
	return ids.formatFunc(func(id ID) (string, error) { return id.FormatOSCALSort() })
}

func (ids IDs) FormatIDType(idType IDType) ([]string, error) {
	switch idType {
	case IDTypeNIST:
		return ids.FormatNIST()
	case IDTypeNISTSort:
		return ids.FormatNISTSort()
	case IDTypeOSCAL:
		return ids.FormatOSCAL()
	case IDTypeOSCALSort:
		return ids.FormatOSCALSort()
	default:
		return []string{}, ErrIDTypeNotSupported
	}
}

func (ids IDs) formatOSCALSortMap() (map[string]int, error) {
	ids2, err := ids.FormatOSCALSort()
	if err != nil {
		return map[string]int{}, err
	}
	out := map[string]int{}
	for _, id := range ids2 {
		out[id]++
	}
	return out, nil
}

func (ids IDs) formatFunc(fn func(id ID) (string, error)) ([]string, error) {
	var out []string
	for _, id := range ids {
		if val, err := fn(id); err != nil {
			return []string{}, err
		} else {
			out = append(out, val)
		}
	}
	out = sliceDedupe(out)
	sort.Strings(out)
	return out, nil
}
