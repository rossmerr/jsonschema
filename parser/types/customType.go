package types

import "github.com/RossMerr/jsonschema/parser/document"

var _ document.Types = (*CustomType)(nil)

// Obsolete
type CustomType struct {
	comment string
	name    string
	Type    string
}

// Obsolete
func NewCustomType(name, typename string) *CustomType {
	return &CustomType{
		name: name,
		Type: typename,
	}
}

func (s *CustomType) WithReference(ref bool) document.Types {
	return s
}

func (s *CustomType) WithFieldTag(tags string) document.Types {
	return s
}

func (s *CustomType) Comment() string {
	return s.comment
}

func (s *CustomType) Name() string {
	return s.name
}

const CustomTypeTemplate = `
{{- define "customtype" -}}
{{ if .Comment -}}
// {{.Comment}}
{{end -}}
 type {{ .Name }} {{ .Type }}
{{end}}
`
