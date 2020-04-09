package jsonschema

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
)

// ID ($id) property is a URI-reference that serves two purposes:
//		It declares a unique identifier for the schema.
//		It declares a base URI against which $ref URI-references are resolved.
type ID string

// NewID returns a new ID
func NewID(s string) (ID, error) {
	uri, err := canonicalURI(s)
	return ID(uri), err
}

// String returns the underlying string representation
func (s ID) String() string {
	return string(s)
}

// Fragment returns the fragment identifier from the ID, everything after the hash mark '#'
func (s ID) Fragment() string {
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

	return basename
}

// canonicalURI is a absolute path to a resource and must be a valid URI
func canonicalURI(s string) (string, error) {
	fragments := strings.Index(s, "#")
	if fragments > 0 {
		s = s[:fragments]
	}

	uri, err := url.Parse(s)
	if err != nil {
		return ".", fmt.Errorf("ID must be a URL '%v' %w", s, err)
	}

	if !uri.IsAbs() {
		return ".", fmt.Errorf("ID's must be a absolute URL '%v'", s)
	}

	if strings.HasSuffix(s, "/") {
		s = s[:len(s)-1]
	}

	return s, nil
}

func (s *ID) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	if v == "" {
		return nil
	}

	_, err = canonicalURI(v)
	if err != nil {
		return err
	}

	*s = ID(v)
	return nil
}
