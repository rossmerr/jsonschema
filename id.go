package jsonschema

import (
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
)

type ID string

func NewID(s string) (ID, error) {
	uri, err := canonicalURI(s)
	return ID(uri), err
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

func (s ID) IsNotEmpty() bool {
	return s != EmptyString
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
