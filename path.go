package jsonschema

type Path []string

func (s Path) Fieldname() string {
	if len(s) > 0 {
		field := s[len(s)-1]
		return Fieldname(field)
	}

	return EmptyString
}
