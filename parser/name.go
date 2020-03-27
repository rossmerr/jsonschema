package parser

import "github.com/RossMerr/jsonschema"

type Name struct {
	fieldname string
	tagname   string
}

func NewName(name string) *Name {
	return &Name{
		fieldname: jsonschema.Structname(name),
		tagname:   name,
	}
}

func (s *Name) Fieldname() string {
	return s.fieldname
}

func (s *Name) Tagname() string {
	return s.tagname
}
