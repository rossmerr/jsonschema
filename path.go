package jsonschema

type Path []string

func (s Path) ToKey() string {
	if len(s) > 0 {
		field := s[len(s)-1]
		return field
	}

	return "."
}
