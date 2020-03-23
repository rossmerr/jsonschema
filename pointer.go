package jsonschema

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

type Pointer string

func NewPointer(s string) Pointer {
	return Pointer(s)
}

func (s Pointer) String() string {
	return string(s)
}

func (s Pointer) Fragments() (query []string) {
	raw := string(s)
	if len(raw) < 1 {
		return []string{}
	}

	uri, err := url.Parse(raw)
	if err != nil {
		log.Print(fmt.Sprintf("Pointer: not a vaild url format '%v'", raw))
		query = []string{}
		return
	}

	if uri.Fragment == EmptyString{
		log.Print(fmt.Sprintf("Pointer: no '#' found in '%v'", raw))
		query = []string{}
		return
	}
	parts := uri.Fragment

	query = strings.Split(parts, "/")
	query = Filter(query, func(v string) bool { return v != "" })
	return
}

// Base reports the file this ID references
func (s Pointer) Base() string {
	raw := string(s)
	if len(raw) < 1 {
		return "."
	}

	index := strings.Index(raw, "#")

	if index >= 0 {
		raw = raw[:index]
	}

	uri, err := url.Parse(raw)
	if err != nil {
		log.Print(fmt.Sprintf("Pointer: not a vaild url format '%v'", raw))
		return "."
	}

	file := filepath.Base(uri.Path)
	return file
}


func (s Pointer) Fieldname() string {
	fragments := s.Fragments()

	if len(fragments) == 0 {
		log.Print(fmt.Sprintf("Pointer: no fragents found in '%v'", string(s)))
		return EmptyString
	}
	name := fragments[len(fragments)-1]

	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}

	clean := reg.ReplaceAllString(name, " ")
	return reg.ReplaceAllString(strings.Title(clean), "")
}
