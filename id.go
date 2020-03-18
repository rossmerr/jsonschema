package jsonschema

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

const definitions = "#/definitions/"

type ID string

func NewID(s string) ID {
	return ID(s)
}

func (s ID) String() string {
	return string(s)
}

func (s ID) pointer() (path string, query []string, err error) {
	raw := string(s)
	if len(raw) < 1 {
		err = fmt.Errorf("No parts found")
		return
	}

	index := strings.Index(raw, "#")
	if index < 0 {
		path = raw
		query = []string{}
		return
	}
	path = raw[:index]
	parts := raw[index+1:]

	query = strings.Split(parts, "/")
	query = Filter(query, func(v string) bool { return v != "" })
	return
}

// Base reports the file this ID references
func (s ID) Base() string {
	path, _, _ := s.pointer()

	return filepath.Base(path)
}

func (s ID) Filename() string {
	basename := s.Base()
	name := strings.TrimSuffix(basename, filepath.Ext(basename))

	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}

	clean := reg.ReplaceAllString(name, " ")
	filename := reg.ReplaceAllString( strings.Title(clean), "")
	return string(unicode.ToLower(rune(filename[0]))) + filename[1:]
}
