package generator

import (
	"strings"
	"encoding/json"
)

func templateStringJoin(rawStrings []string) (string, error) {
	if len(rawStrings) > 0 {
		escapedStrings := []string{}
		for _, rawString := range rawStrings {
			escapedString := strings.Replace(rawString, "$", "$$", -1)
			jsonEscapedString, err := json.Marshal(escapedString)
			if err != nil {
				return "", err
			}
			escapedStrings = append(escapedStrings, string(jsonEscapedString))
		}
		return strings.Join(escapedStrings, ",\n"), nil
	}
	return "", nil
}
