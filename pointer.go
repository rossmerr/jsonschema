package jsonschema

type Pointer []string

func (s Pointer)Fieldname() string {
	if len(s) > 0 {
		field := s[len(s)-1]
		return Fieldname(field)
	}



	return EmptyString
}