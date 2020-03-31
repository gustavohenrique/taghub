package stringutils

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func SliceContains(s []string, term string) bool {
	sort.Strings(s)
	i := sort.SearchStrings(s, term)
	return i < len(s) && s[i] == term
}

func HasBlank(s []string) bool {
	for _, i := range s {
		if IsBlank(i) {
			return true
		}
	}
	return false
}

func IsBlank(s string) bool {
	return strings.Trim(s, " ") == ""
}

func FromJSON(s string, i interface{}) error {
	return json.Unmarshal([]byte(s), &i)
}

func ToJSON(i interface{}) string {
	if i == nil {
		return ""
	}
	b, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(b)
}

func PrettyJSON(i interface{}) string {
	if i == nil {
		return ""
	}
	b, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		return ""
	}
	return string(b)
}

func PrintJSON(sep string, i interface{}) {
	str := PrettyJSON(i)
	fmt.Printf("\n%s\n%s", sep, str)
}

func TrimSpaceNewlineInString(s string) string {
	re := regexp.MustCompile(` +\r?\n +`)
	n := re.ReplaceAllString(s, " ")
	return strings.ReplaceAll(n, "\n", " ")
}
