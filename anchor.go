package jsonschema

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Anchor string

func NewAnchor(s string) Anchor {
	if strings.HasPrefix(s, "#") {
		return Anchor(s)
	}
	return Anchor("#"+s)
}

func (s Anchor) String() string {
	return string(s)
}

func (s Anchor) Fieldname() string {
	raw := string(s)
	index := strings.Index(raw, "#")
	if index < 0 {
		log.Print(fmt.Sprintf("Anchor: no '#' found in '%v'", raw))
		return EmptyString
	}

	name := raw[index+1:]

	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}

	clean := reg.ReplaceAllString(name, " ")
	return reg.ReplaceAllString(strings.Title(clean), "")
}
