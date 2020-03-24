package jsonschema

import (
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

func Contains(a []string, b string) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

func Filter(a []string, delegate func(string) bool) []string {
	list := []string{}
	for _, v := range a {
		if delegate(v) {
			list = append(list, v)
		}
	}

	return list
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

func Fieldname(s string) string {

	name := strings.TrimSuffix(s, filepath.Ext(s))


	// Valid field names must start with a unicode letter
	if !unicode.IsLetter(rune(name[0])) {
		name = "A" + name
	}

	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}

	clean := reg.ReplaceAllString(name, " ")
	return reg.ReplaceAllString( strings.Title(clean), "")
}

func Structname(s string) string {
	return Fieldname(s)
}