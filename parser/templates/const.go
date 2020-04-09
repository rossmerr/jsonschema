package templates

import (
	"github.com/RossMerr/jsonschema/parser"
)

const (
	emptyString = ""
)

var _ parser.Component = (*Const)(nil)

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

func (s *Const) Name() string {
	return emptyString
}

const ConstTemplate = `
{{- define "const" -}}
const (
{{range $key, $item := .List -}}
	{{ typename $item.Type }}_{{ typename $item.Name }} {{ typename $item.Type }} = {{printf "%q" $item.Value }}
{{end -}}
)
{{end -}}
`
