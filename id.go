package jsonschema

import (
	"log"
	"net/url"
	"regexp"
	"strings"
	"unicode"
)

type ID string

func (s ID) Filename() string {
	filename := s.Typename()
	return string(unicode.ToLower(rune(filename[0]))) + filename[1:]
}

func (s ID) Typename() string {
	u, err := url.Parse(string(s))
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
