package tags

import (
	"strings"
)

func KeysString(m map[string]string) string {
	keys := []string{}
	for k, v := range m {
		if v != emptyString {
			keys = append(keys, k+"="+v)
		} else {
			keys = append(keys, k)
		}
	}
	return strings.Join(keys, ",")
}
