package jsonschema

import (
	"fmt"
	"log"
	"net/url"
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

// Deprecated
func (s ID)Title() string {
	return strings.Title(s.String())
}


func (s ID) Pointer() (path string, query []string, err error) {
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
	path, _, _ := s.Pointer()

	return filepath.Base(path)
}

// IsAbs reports whether the URL is absolute.
// Absolute means that it has a non-empty scheme.
func (s ID) IsAbs() bool {
	path, _, err := s.Pointer()
	if err != nil {
		return false
	}

	url, err := url.Parse(path)
	if err != nil {
		return false
	}

	return url.IsAbs()
}



func (s ID) Typename() string {
	basename := s.Base()
	name := strings.TrimSuffix(basename, filepath.Ext(basename))

	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}

	clean := reg.ReplaceAllString(name, " ")
	return reg.ReplaceAllString( strings.Title(clean), "")
}

func (s ID) Filename() string {
	filename := s.Typename()
	return string(unicode.ToLower(rune(filename[0]))) + filename[1:]
}



// Deprecated
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
