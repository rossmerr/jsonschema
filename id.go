package jsonschema

import (
	"net/url"
	"path/filepath"
	"strings"
)

type ID string

func NewID(s string) ID {
	return ID(canonicalURI(s))
}
func (s ID) String() string {
	return string(s)
}


func (s ID) ToTypename() string {
	raw := string(s)
	if len(raw) < 1 {
		return "."
	}

	index := strings.Index(raw, "#")
	var basename string
	if index < 0 {
		basename = filepath.Base(raw)
	} else {
		basename = filepath.Base(raw[:index])
	}


	return ToTypename(basename)
}

func canonicalURI(s string) string {
	fragments := strings.Index(s, "#")
	if fragments > 0 {
		s = s[:fragments]
	}

	uri, err := url.Parse(s)
	if err != nil {
		return "."
	}

	if !uri.IsAbs() {
		return "."
	}

	if strings.HasSuffix(s,"/") {
		s = s[:len(s)-1]
	}

	return s
}