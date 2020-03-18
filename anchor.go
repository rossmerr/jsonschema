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

// func (s Anchor) Fragments() (fragments []string) {
// 	raw := string(s)
// 	if len(raw) < 1 {
// 		return []string{}
// 	}
//
// 	index := strings.Index(raw, "#")
// 	if index < 0 {
// 		log.Print(fmt.Sprintf("Anchor: no '#' found in '%v'", raw))
// 		return []string{}
// 	}
// 	parts := raw[index+1:]
//
// 	fragments = strings.Split(parts, "/")
// 	fragments = Filter(fragments, func(v string) bool { return v != "" })
// 	return
// }

func (s Anchor) Typename() string {
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
