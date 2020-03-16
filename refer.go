package jsonschema

import (
	"log"
	"regexp"
	"strings"
)

type Refer string

func NewRefer(s string) Refer {
	return Refer(s)
}

func (s Refer) String() string {
	return string(s)
}

func (s Refer) Typename() string {
	raw := string(s)
	slashIndex := strings.LastIndex(raw, "/")
	if slashIndex < 0 {
		slashIndex = 0
	}

	path := raw[slashIndex +1 :]
	dotIndex := strings.Index(path, ".")
	if dotIndex < 0 {
		dotIndex = len(path)
	}

	name := path[0:dotIndex]

	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(strings.Title(name), "")
}