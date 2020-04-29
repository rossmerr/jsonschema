package types

import "text/template"

var SchemaFuncMap = template.FuncMap{
}

func SchemaTemplate() (*template.Template, error) {
	tmpl, err := template.New("document").Funcs(SchemaFuncMap).Parse(PkgTemplate)
	if err != nil {
		return nil, err
	}
	return AddTemplates(tmpl)

}

var TemplateArray = []string{
	CommentTemplate,
	StructTemplate,
	InterfaceTemplate,
	FieldTemplate,
	FuncTemplate,
	ParamTemplate,
	PtrTemplate,
	TypeTemplate,
}

func AddTemplates(tmpl *template.Template) (*template.Template, error) {
	var err error
	for _, template := range TemplateArray {
		tmpl, err = tmpl.Parse(template)
		if err != nil {
			return nil, err
		}
	}

	return tmpl, nil
}

const PkgTemplate = `
{{- range $index, $type := .TypeMap -}}
	{{template "type" %type }}
{{- end- }}
`

const CommentTemplate = `
{{- define "struct" -}}
{{ . }}
{{- end -}}
`

const StructTemplate = `
{{- define "struct" -}}
{{if .Comment -}}
// {{.Comment}}
{{end -}}
type {{.Name}} struct {
{{range $index, $field := .Fields -}}
	{{if $field -}}
		{{template "field" $field }}
	{{end -}}
{{end -}}
} {{.Tag}}
{{- end -}}
`

const InterfaceTemplate = `
{{- define "interface" -}}
{{if .Comment -}}
// {{.Comment}}
{{end -}}
type {{.Name}} interface {
{{range $index, $func := .Func -}}
	{{if $field -}}
		{{template "func" $func }}
	{{end -}}
{{end -}}
} {{.Tag}}
{{- end -}}
`

const FieldTemplate = `
{{- define "field" -}}
{{if .Comment -}}
// {{.Comment}}
{{end -}}
{{.Name}} {{.Type}} {{.Tag}}
{{- end -}}
`

const FuncTemplate = `
{{- define "func" -}}
{{if .Comment -}}
// {{.Comment}}
{{end -}}
func {{if .Receiver }}(s *{{.Receiver}}){{end}} {{.Name}}({{range $index, $param := .Params -}}{{template "param" $param }}{{end}}) ({{range $index, $param := .Params -}}{{template "param" $Results }}{{end}}) {{if .Body}} { 
	{{.Body}} 
}
{{end}}
{{- end -}}
`

const ParamTemplate = `
{{- define "param" -}}
{{.Name}} {{.Type}}
{{- end -}}
`

const PtrTemplate = `
{{- define "ptr" -}}
* {{.Elem}}
{{- end -}}
`

const TypeTemplate = `
{{- define "type" -}}

{{- end -}}
`