package jsonschema

import "strings"

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func KeysString(m map[string]string) string {
	keys := make([]string, 0, len(m))
	for k, v := range m {
		if v != EmptyString {
			keys = append(keys, k+"="+v)
		} else {
			keys = append(keys, k)
		}
	}
	return strings.Join(keys, ",")
}
