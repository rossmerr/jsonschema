package types

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
)

var _ document.Types = (*Root)(nil)

type Root struct {
	comment string
	Type    document.Types
	Methods []string
}

func HandleRoot(ctx *document.DocumentContext, name string, schema *jsonschema.Schema) (document.Types, error) {
	return NewRoot(ctx, name, schema)
}

func NewRoot(ctx *document.DocumentContext, name string, schema *jsonschema.Schema) (document.Types, error) {
	t, err := ctx.Process(name, schema)
	return &Root{
		comment: schema.Description,
		Type:    t,
	}, err
}

func (s *Root) WithReference(ref bool) document.Types {
	return s
}

func (s *Root) WithFieldTag(tags string) document.Types {
	return s
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
func (s {{ $.Name }}) {{$method}}(){}
{{end }}
{{- end -}}
`
