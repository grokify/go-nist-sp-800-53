package rev5

import "strings"

type FamilySet struct {
	Data map[string]string
}

func NewFamilySet() FamilySet {
	return FamilySet{Data: familiesMap()}
}

func IsControlFamilyAbbreviation(abbr string) bool {
	abbr = strings.ToUpper(strings.TrimSpace(abbr))
	m := familiesMap()
	_, ok := m[abbr]
	return ok
}

func familiesMap() map[string]string {
	return map[string]string{
		"AC": "Access Control",
		"AT": "Awareness and Training",
		"AU": "Audit and Accountability",
		"CA": "Assessment, Authorization, and Monitoring",
		"CM": "Configuration Management",
		"CP": "Contingency Planning",
		"IA": "Identification and Authentication",
		"IR": "Incident Response",
		"MA": "Maintenance",
		"MP": "Media Protection",
		"PE": "Physical and Environmental Protection",
		"PL": "Planning",
		"PM": "Program Management",
		"PS": "Personnel Security",
		"PT": "Personally Identifiable Information Processing and Transparency",
		"RA": "Risk Assessment",
		"SA": "System and Services Acquisition",
		"SC": "System and Communications Protection",
		"SI": "System and Information Integrity",
		"SR": "Supply Chain Risk Management",
	}
}
