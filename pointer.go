package jsonschema

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
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

