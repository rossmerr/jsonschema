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

	// tmpl, err = AddTemplates(tmpl, err)
	// if err != nil {
	// 	return nil, err
	// }

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

const UnmarshalJSONBodyTemplate = `
m := map[string]json.RawMessage{}
if err := json.Unmarshal(b, &m); err != nil {
	return nil
}

{{range $key, $ref := .References -}}
{{ mixedCase $ref.Name }} := func() {{ $ref.Type }} {
	raw, ok := m["{{ $ref.Name }}"]
	if !ok {
	return nil
	}
	var {{ mixedCase $ref.Name }} {{ $ref.Type }}
	if err := json.Unmarshal(raw, &{{ mixedCase $ref.Name }}); err != nil {
		return nil
	}
	return {{ mixedCase $ref.Name }}
}
{{end}}
type Alias {{.Type}}
aux := &struct {
{{range $key, $ref := .References -}}
	{{ $ref.Name }} {{ $ref.FieldTag }}
{{end -}}
*Alias
}{
	{{range $key, $ref := .References -}}
		{{ $ref.Name }}: {{ mixedCase $ref.Name }},
	{{end -}}
	Alias: (*Alias)(s),
}

if err := json.Unmarshal(b, &aux); err != nil {
	return err
}

{{range $key, $ref := .References -}}
s.{{ $ref.Name }} = aux.{{ $ref.Name }}
{{end}}
return nil`
