package types

import (
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
	"github.com/RossMerr/jsonschema/parser/tags"
	"github.com/RossMerr/jsonschema/parser/tags/json"
	"github.com/RossMerr/jsonschema/parser/tags/validate"
)

var _ document.Types = (*Struct)(nil)

type Struct struct {
	comment   string
	name      string
	Fields    []document.Types
	StructTag string
	FieldTag  string
}

func HandleObject(ctx *document.DocumentContext, name string, schema *jsonschema.Schema) (document.Types, error) {
	return NewStruct(ctx, name, schema)
}

func NewStruct(ctx *document.DocumentContext, name string, schema *jsonschema.Schema) (document.Types, error) {

	fields := []document.Types{}
	for key, propertie := range schema.Properties {
		s, err := ctx.Process(key, propertie)
		if err != nil {
			return nil, err
		}

		tags := tags.NewFieldTag([]tags.StructTag{json.NewJSONTags(), validate.NewValidateTags()})
		fieldTag := tags.ToFieldTag(key, propertie, schema.Required)

		ref := !jsonschema.Contains(schema.Required, strings.ToLower(key))

		s.WithFieldTag(fieldTag).WithReference(ref)

		if enum, ok := s.(*Enum); ok {
			fields = append(fields, GlobalEnum(ctx, enum, name))
			continue
		}

		fields = append(fields, s)
	}

	for key, def := range schema.AllDefinitions() {
		s, err := NewRoot(ctx, key, def)
		if err != nil {
			return nil, err
		}

		ctx.Globals[key] = s
	}

	return &Struct{
		comment: schema.Description,
		name:    jsonschema.ToTypename(name),
		Fields:  fields,
	}, nil
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
