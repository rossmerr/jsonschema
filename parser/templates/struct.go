package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Receiver = (*Struct)(nil)

type Struct struct {
	comment  string
	name     string
	Fields   []parser.Component
	fieldTag string
	Reference string
	Methods  []*parser.Method
}

func NewStruct(name, comment string, fields ...parser.Component) *Struct {
	return &Struct{
		comment: comment,
		name:    name,
		Fields:  fields,
	}
}

func (s *Struct) unmarshalStructJSON() *parser.Method {
	var references []*Reference
	for _, field := range s.Fields {
		switch f := field.(type) {
		case *OneOf:
			references = append(references, f.Reference)
		case *AnyOf:
			references = append(references, f.Reference)
		}
	}

	if len(references) == 0 {
		return nil
	}

	unmarshal, err := MethodUnmarshalJSON(s.Name(), references)
	if err != nil {
		panic(err)
	}

	return unmarshal
}

func (s *Struct) WithMethods(methods ...*parser.Method) {
	s.Methods = append(s.Methods, methods...)
}

func (s *Struct) WithReference(ref bool) parser.Field {
	if ref {
		s.Reference = "*"
	} else {
		s.Reference = ""
	}
	return s
}

func (s *Struct) WithFieldTag(tags string) parser.Field {
	s.fieldTag = tags
	return s
}

func (s *Struct) FieldTag() string {
	return s.fieldTag
}

func (s *Struct) Comment() string {
	return s.comment
}

func (s *Struct) Name() string {
	return s.name
}

const StructTemplate = `
{{- define "struct" -}}
{{ typename .Name }} {{ .Reference}}struct {
{{range $key, $propertie := .Fields -}}
	{{template "kind" $propertie }}
{{end -}}
} {{ .FieldTag }}

{{range $key, $method := .Methods -}}
	{{ if $method }}
		{{template "method" $method }}
	{{ end}}
{{end }}
{{- end -}}`
