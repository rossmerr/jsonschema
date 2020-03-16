package jsonschema

import (
	"log"
	"net/url"
	"regexp"
	"strings"
	"unicode"
)

const definitions = "#/definitions/"

type ID string

func NewID(s string) ID {
	return ID(s)
}

func NewDefinitionsID(s ID) ID {
	return NewID(definitions + s.String())
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

func (s ID) Parts() (schema string, fragment string, typename string, err error) {
	raw := string(s)
	if len(raw) < 1 {
		//err = fmt.Errorf("No parts found")
		return
	}

	uri, err := url.Parse(raw)
	if err != nil {
		return
	}

	var name string


	if uri.Fragment != "" {
		index := strings.Index(raw, uri.Path)
		schema = raw[0:index] + uri.Path
		parts := strings.SplitAfter(uri.Fragment, "/")
		name = parts[len(parts)-1]
		fragment = strings.Join(parts[0:len(parts)-1], "")
	} else {
		name = uri.Path
	}





	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}
	typename = reg.ReplaceAllString(strings.Title(name), "")
	return
}

func (s ID) Typename() (string) {
	_, _, typename, err := s.Parts()
	if err != nil {
		log.Fatal(err)
	}
	return typename
}
