package jsonschema

import "strings"

func Contains(a []string, b string) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

func ForEach(a []string, delegate func(string) string) []string {
	list := []string{}
	for _, v := range a {
		list = append(list, delegate(v))
	}

	return list
}

func Unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, v := range slice {
		if _, value := keys[v]; !value {
			keys[v] = true
			list = append(list, v)
		}
	}
	return list
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
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
