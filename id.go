package jsonschema

import (
	"log"
	"regexp"
	"strings"
	"unicode"
)

type ID string

func NewID(s string) ID {
	return ID(s)
}

func NewDefinitionsID(s ID) ID {
	return NewID("#/definitions/" + s.String())
}

func (s ID) String() string {
	return string(s)
}


func (s ID)Title() string {
	return strings.Title(s.String())
}
func (s ID) Filename() string {
	filename := s.Typename()
	return string(unicode.ToLower(rune(filename[0]))) + filename[1:]
}

func (s ID) Typename() string {

	raw := string(s)
	if len(raw) < 1 {
		return raw
	}

	slashIndex := strings.LastIndex(raw, "/")


	path := raw[slashIndex +1 :]
	dotIndex := strings.Index(path, ".")
	if dotIndex < 0 {
		dotIndex = len(path)
	}

	name := path[0:dotIndex]

	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(strings.Title(name), "")
}
