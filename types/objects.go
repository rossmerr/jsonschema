package types

import (
	"encoding/json"
	"reflect"

	"github.com/RossMerr/jsonschema/tokens"
)

const (
	Object = "object"
)

type Objects struct {
	T             string          `json:"Type,omitempty"`
	Description   string          `json:"Description,omitempty"`
	MaxProperties *uint32         `json:"MaxProperties,omitempty"`
	MinProperties *uint32         `json:"MinProperties,omitempty"`
	Required      []string        `json:"Required,omitempty"`
	OneOf         OneOf           `json:"oneOf"`
	Properties    map[string]Type `json:"-"`
}

//
// func (s Objects) InterfaceTypename() string {
// 	u, err := url.Parse(s.ID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	path := strings.Trim(u.Path, "/")
// 	index := strings.Index(path, ".")
// 	if index < 0 {
// 		index = len(path)
// 	}
//
// 	reg, err := regexp.Compile(`[^a-zA-Z0-9]+`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return reg.ReplaceAllString(strings.Title(path[0:index]), "")
// }


func (s Objects) Type() reflect.Kind {
	if s.T == Object {
		if s.OneOf != nil && len(s.OneOf) > 0 {
			return reflect.Interface
		}
		return reflect.Struct

	}

	return reflect.Invalid
}

func (s *Objects) UnmarshalJSON(data []byte) error {

	type Alias Objects
	err := json.Unmarshal(data, &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	})

	properties, err := PropertiesUnmarshalJSON(data, tokens.Properties)
	if err != nil {
		return err
	}
	s.Properties = properties

	return err
}
