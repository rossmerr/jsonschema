package types

type Type interface {
	t()
}

type Comment string

type Struct struct {
	Name string
	Comment Comment
	Tag StructTag
	Fields Fields
	pkg *Pkg
}

func (s *Struct) t() {}

type StructTag string

type Field struct {
	Name string
	Comment Comment
	Type *Type
	Tag FieldTag
}

type Fields []*Field

type FieldTag string

type Ptr struct {
	Elem *Type
}

func (s *Ptr) t() {}

type Interface struct {
	Funcs []*Func
	pkg *Pkg
}

func (s *Interface) t() {}

type Tuple struct {
	first  *Type
	second *Type
}

func (s *Tuple) t() {}

type Array struct {
	Elem  *Type // element type
}

func (s *Array) t() {}

type Func struct {
	Name string
	Comment Comment
	Receiver *Type
	Results  []*Type
	Params   []*Type
	Body string
	pkg *Pkg
}

func (s *Func) t() {}