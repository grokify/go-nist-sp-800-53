package rev5

import (
	"slices"
	"sort"
)

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
