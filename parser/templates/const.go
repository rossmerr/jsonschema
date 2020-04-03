package templates

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*Const)(nil)

type Const struct {
	List []*ConstItem
}

type ConstItem struct {
	Name  string
	Type  string
	Value interface{}
}

func NewConst(list ...*ConstItem) *Const {
	return &Const{
		List: list,
	}
}

func (s *Const) WithMethods(methods ...*parser.Method) parser.Types {
	return s
}

func (s *Const) WithReference(ref bool) parser.Types {
	return s
}

func (s *Const) WithFieldTag(tags string) parser.Types {
	return s
}

func (s *Const) FieldTag() string {
	return jsonschema.EmptyString
}

func (s *Const) Comment() string {
	return jsonschema.EmptyString
}

func (s *Const) Name() string {
	return jsonschema.EmptyString
}

const ConstTemplate = `
{{- define "const" -}}
const (
{{range $key, $item := .List -}}
	{{ $item.Type }}_{{ title $item.Name }} {{ $item.Type }} = {{printf "%q" $item.Value }}
{{end -}}
)
{{end -}}
`
