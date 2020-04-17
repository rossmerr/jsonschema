package templates

import (
	"bytes"
	"text/template"

	"github.com/RossMerr/jsonschema/parser"
)

func MethodUnmarshalJSON(receiver string, references []*Reference) (*parser.Method, error) {
	methodSignature := parser.NewMethodSignature("UnmarshalJSON")
	methodSignature.WithInputs(parser.NewParameter("b", "[]byte"))
	methodSignature.WithOutputs(parser.NewParameter("", "error"))
	method := parser.NewMethodFromSignature(receiver, methodSignature)

	tmpl, err := template.New("method").Funcs(SchemaFuncMap).Parse(UnmarshalJSONBodyTemplate)
	if err != nil {
		return nil, err
	}

	tmpl, err = tmpl.Parse(arrayTemplate)
	if err != nil {
		return nil, err
	}

	tmpl, err = tmpl.Parse(objectTemplate)
	if err != nil {
		return nil, err
	}

	var byte bytes.Buffer
	err = tmpl.Execute(&byte, &struct {
		References []*Reference
		Type       string
	}{
		References: references,
		Type:       receiver,
	})
	if err != nil {
		return nil, err
	}

	method.Body = byte.String()
	return method, nil

}

type Test struct{}

const arrayTemplate = `
{{- define "arrayunmarshaljson" -}}
{{ mixedCase .Name }} := func() []{{typename .Type.Name }} {
	var {{ mixedCase .Name }} []{{typename .Type.Name }}
	raw, ok := m["{{ .Name }}"]
	if !ok {
		return storage
	}

	var array []json.RawMessage
	if err := json.Unmarshal(raw, &array); err != nil {
		return storage
	}

	for _, item := range array {
	{{- range $key, $type := .Types }}
		var {{ mixedCase $type}} *{{typename $type}}
		if err := json.Unmarshal(item, &{{ mixedCase $type}}); err == nil {
			{{ mixedCase $.Name }} = append({{ mixedCase $.Name }}, {{ mixedCase $type}})
		}
	{{end -}}
	}

	return {{ mixedCase .Name }}
}

{{- end -}}
`

const objectTemplate = `
{{- define "objectunmarshaljson" -}}
{{ mixedCase .Name }} := func() {{typename .Type.Name }} {
	raw, ok := m["{{ .Name }}"]
	if !ok {
	return nil
	}
	{{range $key, $type := .Types }}
		var {{ mixedCase $type}} {{typename $type}}
		if err := json.Unmarshal(raw, &{{ mixedCase $type}}); err == nil {
			return &{{ mixedCase $type}}
		}
	{{end}}
		
	return nil
}
{{- end -}}
`

const UnmarshalJSONBodyTemplate = `
m := map[string]json.RawMessage{}
if err := json.Unmarshal(b, &m); err != nil {
	return nil
}

{{range $key, $ref := .References -}}
	{{if eq $ref.Type.Kind.String "array"}}
		{{- template "arrayunmarshaljson" $ref -}}
	{{else}}
		{{- template "objectunmarshaljson" $ref -}}
	{{end -}}

{{end}}
type Alias {{typename .Type}}
aux := &struct {
{{range $key, $ref := .References -}}
	{{if eq $ref.Type.Kind.String "array"}}
		{{typename $ref.Name }} []{{typename $ref.Type.Name }}{{ $ref.FieldTag }}
	{{else}}
		{{typename $ref.Name }} {{typename $ref.Type.Name }}{{ $ref.FieldTag }}
	{{end -}}
{{end -}}
*Alias
}{
	{{range $key, $ref := .References -}}
		{{typename $ref.Name }}: {{ mixedCase $ref.Name }}(),
	{{end -}}
	Alias: (*Alias)(s),
}

{{range $key, $ref := .References -}}
s.{{typename $ref.Name }} = aux.{{typename $ref.Name }}
{{end}}
return nil`
