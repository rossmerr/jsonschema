package templates

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*Root)(nil)

type Root struct {
	comment string
	Type    parser.Types
	Methods []*parser.Method
}

func NewRoot(comment string, t parser.Types) parser.Types {
	root := &Root{
		comment: comment,
		Type:    t,
		Methods: []*parser.Method{},
	}

	switch s := t.(type) {
	case *Struct:
		root.unmarshalStructJSON(s)
	}

	return root
}

func (s *Root) unmarshalStructJSON(str *Struct) {
	references := []*Reference{}
	for _, field := range str.Fields {
		switch f := field.(type) {
		case *OneOf:
			references = append(references, f.Reference)
		case *AnyOf:
			references = append(references, f.Reference)
		case *AllOf:
			references = append(references, f.Reference)
		}
	}

	if len(references) == 0 {
		return
	}

	unmarshal, err := MethodUnmarshalJSON(str.Name(), references)
	if err != nil {
		panic(err)
	}
	if unmarshal != nil {
		s.Methods = append(s.Methods, unmarshal)
	}
}

func (s *Root) WithMethods(methods ...*parser.Method) parser.Types {
	s.Methods = append(s.Methods, methods...)
	return s
}

func (s *Root) WithReference(ref bool) parser.Types {
	return s
}

func (s *Root) WithFieldTag(tags string) parser.Types {
	return s
}

func (s *Root) FieldTag() string {
	return jsonschema.EmptyString
}

func (s *Root) Comment() string {
	return s.comment
}

func (s *Root) Name() string {
	return s.Type.Name()
}

const RootTemplate = `
{{- define "root" -}}
{{ if .Type.Comment -}}
// {{.Type.Comment}}
{{end -}}
type {{template "kind" .Type }}

{{range $key, $method := .Methods -}}
	{{template "method" $method }}
{{end }}
{{- end -}}
`
