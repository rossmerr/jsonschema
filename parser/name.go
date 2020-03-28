package parser

import "github.com/RossMerr/jsonschema"

type Name struct {
	fieldname string
	tagname   string
}

func NewName(name string) *Name {
	return &Name{
		fieldname: jsonschema.ToTypename(name),
		tagname:   name,
	}
}

func NameFromID(id jsonschema.ID) *Name {
	return &Name{
		fieldname: id.ToTypename(),
		tagname:   id.ToTypename(),
	}
}

func (s *Name) Fieldname() string {
	return s.fieldname
}

func (s *Name) Tagname() string {
	return s.tagname
}
