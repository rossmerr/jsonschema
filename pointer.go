package jsonschema

type Pointer string

func (s Pointer) String() string {
	return string(s)
}
