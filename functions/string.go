package functions

import (
	"encoding/json"
	"log"
	"net/url"
	"regexp"
	"strings"
	"unicode"
)

func ToString (raw json.RawMessage) string {
	var s string
	err := json.Unmarshal(raw, &s)
	if err != nil {
		panic(err)
	}
	return s
}

func TitleCase(raw string) string {
	return strings.Title(raw)
}

func MixedCase(raw string) string {
	s := strings.Replace(strings.Title(raw), " ", "", -1)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func Filename(raw string) string {
	filename := Typename(raw)
	return string(unicode.ToLower(rune(filename[0]))) + filename[1:]
}

func Typename(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		log.Fatal(err)
	}

	path := strings.Trim(u.Path, "/")
	index := strings.Index(path, ".")
	if index < 0 {
		index = len(path)
	}

	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(strings.Title(path[0:index]), "")
}
