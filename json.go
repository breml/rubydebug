package rubydebug

import "regexp"

// ToJSON converts any given logstash rubydebug string
// to valid JSON
func ToJSON(rubydebug string) string {
	doubleArrow := regexp.MustCompile(`(?m)^(\s+\"[^"]+\") => `)
	arrayIndex := regexp.MustCompile(`(?m)^(\s+)\[[0-9]+\]`)
	json := doubleArrow.ReplaceAllString(rubydebug, "${1}: ")
	json = arrayIndex.ReplaceAllString(json, "${1}")
	return json
}
