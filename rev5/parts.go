package rev5

import (
	"strings"

	oscalTypes "github.com/defenseunicorns/go-oscal/src/types/oscal-1-1-3"
)

type Parts []oscalTypes.Part

func (parts Parts) ExtractProseString() string {
	return strings.TrimSpace(strings.Join(parts.ExtractProseLines(), SepLines))
}

func (parts Parts) ExtractProseLines() []string {
	var proseList []string
	for _, part := range parts {
		if part.Prose != "" {
			proseList = append(proseList, part.Prose)
		}
		// Recursively extract prose from sub-parts
		if part.Parts != nil && len(*part.Parts) > 0 {
			proseList = append(proseList, Parts(*part.Parts).ExtractProseString())
		}
	}
	return proseList
}
