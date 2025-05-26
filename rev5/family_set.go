package rev5

import (
	"fmt"
	"strings"
)

type FamilySet struct {
	ControlIDs *ControlIDs
	Data       map[string]Family
}

func NewFamilySet(populate bool) *FamilySet {
	out := FamilySet{
		ControlIDs: NewControlIDs(),
		Data:       map[string]Family{},
	}
	if populate {
		out.Data = familiesMapAll()
	}
	return &out
}

func (set *FamilySet) AddControlID(idAny string, updateFamilyCounts bool) error {
	ido, err := ParseID(idAny)
	if err != nil {
		return err
	}
	var fam Family
	if famTry, ok := set.Data[ido.FamilyAbbreviation]; ok {
		fam = famTry
	} else {
		fam = Family{}
	}
	if err := fam.AddControlID(set.ControlIDs, idAny, updateFamilyCounts); err != nil {
		return err
	}
	set.Data[ido.FamilyAbbreviation] = fam
	return nil
}

func (set *FamilySet) PopulateDefault() {
	if set.Data == nil {
		set.Data = map[string]Family{}
	}
	mapAll := familiesMapAll()
	for k, v := range mapAll {
		if _, ok := set.Data[k]; !ok {
			set.Data[k] = v
		}
	}
}

func (set *FamilySet) UpdateCounts() {
	for abbr, fam := range set.Data {
		fam.UpdateCounts()
		set.Data[abbr] = fam
	}
}

func IsFamilyAbbreviation(abbr string) bool {
	abbr = strings.ToLower(abbr)
	m := familiesMapAllSimple()
	_, ok := m[abbr]
	return ok
}

func familiesMapAll() map[string]Family {
	raw := familiesMapAllSimple()
	out := map[string]Family{}
	for k, v := range raw {
		out[k] = Family{
			Abbreviation: k,
			Name:         v,
		}
	}
	return out
}

func familiesMapAllSimple() map[string]string {
	return map[string]string{
		"ac": "Access Control",
		"at": "Awareness and Training",
		"au": "Audit and Accountability",
		"ca": "Assessment, Authorization, and Monitoring",
		"cm": "Configuration Management",
		"cp": "Contingency Planning",
		"ia": "Identification and Authentication",
		"ir": "Incident Response",
		"ma": "Maintenance",
		"mp": "Media Protection",
		"pe": "Physical and Environmental Protection",
		"pl": "Planning",
		"pm": "Program Management",
		"ps": "Personnel Security",
		"pt": "Personally Identifiable Information Processing and Transparency",
		"ra": "Risk Assessment",
		"sa": "System and Services Acquisition",
		"sc": "System and Communications Protection",
		"si": "System and Information Integrity",
		"sr": "Supply Chain Risk Management",
	}
}

func FamilyAbbrAndNameCanonical(s string) (string, string, error) {
	m := familiesMapAllSimple()
	tryLowerAny := strings.ToLower(strings.TrimSpace(s))
	if v, ok := m[tryLowerAny]; ok {
		return tryLowerAny, v, nil
	}
	for k, v := range m {
		if tryLowerAny == strings.ToLower(v) {
			return k, v, nil
		}
	}
	return "", "", fmt.Errorf("family abbreviation and name not found (%s)", s)
}
