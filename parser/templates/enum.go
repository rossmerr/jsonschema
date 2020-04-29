package templates

import (
	"bytes"
	"text/template"

	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Receiver = (*Enum)(nil)

type Enum struct {
	comment   string
	name      string
	Type      string
	Values    []string
	fieldTag  string
	Reference string
	Items     []*ConstItem
	Methods   []*parser.Method
}

func NewEnum(name, comment, typename string, values []string, items []*ConstItem) *Enum {
	enum := &Enum{
		comment: comment,
		name:    name,
		Type:    typename,
		Values:  values,
		Items:   items,
		Methods: []*parser.Method{},
	}

	unmarshal, err := enumUnmarshalJSON(name, enum)
	if err != nil {
		panic(err)
	}

	enum.Methods = append(enum.Methods, unmarshal)

	return enum
}

func (s *Enum) WithReference(ref bool) parser.Field {
	return s
}

func (s *Enum) WithFieldTag(tags string) parser.Field {
	s.fieldTag = tags
	return s
}

func (s *Enum) WithMethods(methods ...*parser.Method) {
	s.Methods = append(s.Methods, methods...)
}

func (s *Enum) FieldTag() string {
	return s.fieldTag
}

func (s *Enum) Comment() string {
	return s.comment
}

func (s *Enum) Name() string {
	return s.name
}

const EnumTemplate = `
{{- define "enum" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
type {{ typename .Name}} {{ .Type }} 

{{range $key, $method := .Methods -}}
	{{ if $method }}
		{{template "method" $method }}
	{{ end}}
{{end }}
{{end -}}
`

func enumUnmarshalJSON(receiver string, enum *Enum) (*parser.Method, error) {
	methodSignature := parser.NewMethodSignature("UnmarshalJSON")
	methodSignature.WithInputs(parser.NewParameter("b", "[]byte"))
	methodSignature.WithOutputs(parser.NewParameter("", "error"))
	method := parser.NewMethodFromSignature(receiver, methodSignature)

	tmpl, err := template.New("method").Funcs(SchemaFuncMap).Parse(unmarshalEnumJSONBodyTemplate)
	if err != nil {
		return nil, err
	}

	var byte bytes.Buffer
	err = tmpl.Execute(&byte, &struct {
		Enum *Enum
		Type string
	}{
		Enum: enum,
		Type: receiver,
	})
	if err != nil {
		return nil, err
	}

	method.Body = byte.String()
	return method, nil

}

const unmarshalEnumJSONBodyTemplate = `
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	*s = {{ typename .Enum.Name }}(v)

	switch *s {
	{{range $key, $item := .Enum.Items -}}
		case {{ typename $item.Type }}_{{ typename $item.Name }}:
			return nil
	{{end}}}
return fmt.Errorf("{{ typename .Enum.Name }}: value '%v' does not match any value", v)`
