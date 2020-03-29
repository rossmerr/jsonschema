package jsonschema

type Path []string

func (s Path) ToFieldname() string {
	if len(s) > 0 {
		field := s[len(s)-1]
		return ToTypename(field)
	}

	return "."
}
