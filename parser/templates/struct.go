package templates

import (
	"fmt"
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*Struct)(nil)

type Struct struct {
	comment  string
	name     string
	Fields   []parser.Types
	fieldTag string
}

func NewStruct(name, comment string, fields []parser.Types) *Struct {
	return &Struct{
		comment: comment,
		name:    jsonschema.ToTypename(name),
		Fields:  fields,
	}
}

func (s *Struct) WithMethods(methods ...*parser.Method) parser.Types {
	return s
}

func (s *Struct) WithReference(ref bool) parser.Types {
	return s
}

func (s *Struct) WithFieldTag(tags string) parser.Types {
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

func (s *Struct) IsNotEmpty() bool {
	return len(s.Fields) > 0
}

func (s *Struct) UnmarshalJSON() *parser.Method {

	overrideFields := []string{}
	for _, field := range s.Fields {
		switch f := field.(type) {
		case *OneOf:
			name := f.Reference.Name()
			tag := f.Reference.FieldTag()
			typename := f.Reference.Type
			overrideFields = append(overrideFields, name+" "+typename+" "+tag)
		case *AnyOf:
			name := f.Reference.Name()
			tag := f.Reference.FieldTag()
			typename := f.Reference.Type

			overrideFields = append(overrideFields, name+" "+typename+" "+tag)
		case *AllOf:
			name := f.Struct.Name()
			tag := f.Struct.FieldTag()
			typename := "struct"

			overrideFields = append(overrideFields, name+" "+typename+" "+tag)
		}
	}

	if len(overrideFields) == 0 {
		return nil
	}

	method := parser.NewMethod(s.Name(), "UnmarshalJSON")
	method.WithInputs(parser.NewParameter("b", "[]byte"))
	method.WithOutputs(parser.NewParameter("", "error"))
	t := `type Alias %v
	aux := &struct {
		%v
		*Alias
	}{
		
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	return nil`
	method.Body = fmt.Sprintf(t, s.Name(), strings.Join(overrideFields, "\n"))
	return method
}

const StructTemplate = `
{{- define "struct" -}}
{{ .Name }} struct {
{{range $key, $propertie := .Fields -}}
	{{template "kind" $propertie }}
{{end -}}
} {{ .FieldTag }}
{{- end -}}`
