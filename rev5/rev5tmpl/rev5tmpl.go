package rev5tmpl

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const TemplateVarInsertPrefix = "{{ insert:"

// tmplRegex matches {{ insert: param, id }} pattern
var tmplRegex = regexp.MustCompile(`{{\s*insert:\s*(\w+)\s*,\s*([\w\-\.]+)\s*}}`)

// var rxTemplateParam = regexp.MustCompile(`\{\{\sinsert:\sparam,\s(\S+)\s\}\}`)

// Render replaces OSCAL-style template variables with provided values, such as the following:
// `{{ insert: param, <param-id> }}`,
// `{{ insert: component-description, <comp-id> }}`,
// `{{ insert: implementation-status, <control-id> }}`,
// `{{ insert: control-description, <control-id> }}`.
func Render(input string, paramMap map[string]string) (string, error) {
	paramMap, err := renderMap(paramMap)
	if err != nil {
		return "", err
	} else {
		return renderTemplate(input, paramMap)
	}
}

func renderTemplate(input string, paramMap map[string]string) (string, error) {
	unresolvedIDs := map[string]int{}
	unsupportedKinds := map[string]int{}
	rendered := tmplRegex.ReplaceAllStringFunc(input, func(match string) string {
		parts := tmplRegex.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match // malformed
		}

		kind := strings.ToLower(parts[1])
		id := strings.TrimSpace(parts[2])

		switch kind {
		case "param":
			if val, ok := paramMap[id]; ok {
				return val
			} else {
				unresolvedIDs[id]++
				return "[UNRESOLVED:" + id + "]"
			}
		default:
			unsupportedKinds[kind]++
			return "[UNSUPPORTED:" + kind + "]"
		}
	})
	if strings.Contains(rendered, TemplateVarInsertPrefix) {
		return "", errors.New("rev5tmpl: not fully rendered")
	} else if len(unresolvedIDs) > 0 {
		return "", fmt.Errorf("unresolved ids count (%d) ids (%v)", len(unresolvedIDs), unresolvedIDs)
	} else if len(unsupportedKinds) > 0 {
		return "", fmt.Errorf("unsupported kinds count (%d) kinds (%v)", len(unsupportedKinds), unsupportedKinds)
	} else {
		return rendered, nil
	}
}

// renderMap renders template variables in the param map as the values
// may contain variables as well.
func renderMap(paramMap map[string]string) (map[string]string, error) {
	for k, v := range paramMap {
		if strings.Contains(v, TemplateVarInsertPrefix) {
			if vrendered, err := renderTemplate(v, paramMap); err != nil {
				return map[string]string{}, err
			} else {
				paramMap[k] = vrendered
			}
		}
	}
	return paramMap, nil
}
