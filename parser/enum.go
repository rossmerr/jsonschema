package parser

type Enum struct {
	comment   string
	name      string
	Type      string
	Values    []string
	FieldTag  string
	Reference string
}

func NewEnum(ctx *SchemaContext, name *Name, description, fieldTag string, isReference bool, values []string) *Enum {
	reference := ""
	if isReference {
		reference = "*"
	}
	parent := ctx.Parent()

	typename := parent.ID.ToTypename() + name.Fieldname()

//	list := List{
		//NewCustomType(typename, "string"),
//	}

	// for _, value := range values  {
	// 	c := NewConst(strings.Title(value), typename, value)
	// 	list = append(list, c)
	// }

	//ctx.Globals[typename] = &list

	return &Enum{
		comment:   description,
		name:      name.Fieldname(),
		Type:      typename,
		FieldTag:  fieldTag,
		Reference: reference,
		Values:    values,
	}
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
{{ .Name}} {{ .Type }} {{ .FieldTag }}
{{end -}}
`