package parser

type Types interface {
	Comment()     string
	Name() string
}

type Interfaces interface {
	AppendMethods([]string)
}
