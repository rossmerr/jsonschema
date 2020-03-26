package jsonschema

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

type Reference string

func NewReference(s string) Reference {
	return Reference(s)
}

func (s Reference) String() string {
	return string(s)
}

func (s Reference) Pointer() (pointer Pointer) {
	raw := string(s)
	if len(raw) < 1 {
		return Pointer{}
	}

	uri, err := url.Parse(raw)
	if err != nil {
		return Pointer{}
	}

	if uri.Fragment == EmptyString{
		return Pointer{}
	}
	parts := uri.Fragment

	query := strings.Split(parts, "/")
	query = Filter(query, func(v string) bool { return v != "" })
	return query
}

// Base reports the file this ID references
func (s Reference) Base() string {
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
		log.Print(fmt.Sprintf("Reference: not a vaild url format '%v'", raw))
		return "."
	}

	file := filepath.Base(uri.Path)
	return file
}


func (s Reference) Fieldname() string {
	fragments := s.Pointer()

	if len(fragments) == 0 {
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

func (s Reference) IsEmpty() bool {
	return s == EmptyString
}