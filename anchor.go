package jsonschema

import (
	"strings"
)

// Anchor
type Anchor string

func NewAnchor(s string) Anchor {
	return Anchor(s)
}

func (s Anchor) String() string {
	return string(s)
}

func (s Anchor) Fieldname() string {
	raw := string(s)
	index := strings.Index(raw, "#")
	if index < 0 {
		return EmptyString
	}

	return Title(raw[index+1:])
}
