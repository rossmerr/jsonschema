package types

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
)

var _ document.Types = (*Struct)(nil)

type Struct struct {
	comment   string
	name      string
	Fields    []document.Types
	StructTag string
	FieldTag  string
}

func NewStruct(name, comment string, fields []document.Types) document.Types {
	return &Struct{
		comment: comment,
		name:    jsonschema.ToTypename(name),
		Fields:  fields,
	}
}

func (s *Struct) WithReference(ref bool) document.Types {
	return s
}

func (s *Struct) WithFieldTag(tags string) document.Types {
	s.FieldTag = tags
	return s
}

func (s *Struct) Comment() string {
	return s.comment
}

func (s *Struct) Name() string {
	return s.name
}

func (s *Struct) IsNotEmpty() bool {
	return len(s.Fields) > 0
}

const StructTemplate = `
{{- define "struct" -}}
{{ .Name }} struct {
{{range $key, $propertie := .Fields -}}
	{{template "kind" $propertie }}
{{end -}}
} {{ .StructTag }}
{{- end -}}`
