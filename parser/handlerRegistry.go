package parser

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

var (
	ErrNotRegistred = errors.New("This Handler is not registered.")
)

type HandlerRegistry struct {
	handlers map[Kind]HandleSchemaFunc
}

func NewHandlerRegistry() *HandlerRegistry {
	return &HandlerRegistry{
		handlers: map[Kind]HandleSchemaFunc{},
	}
}

func (s HandlerRegistry) RegisterHandler(kind Kind, handler HandleSchemaFunc) {
	if _, found := s.handlers[kind]; found {
		log.Panicf("Already registered Handler %q.", kind)
	}
	s.handlers[kind] = handler
}

func (s HandlerRegistry) ResolveHandler(kind Kind) HandleSchemaFunc {
	if v, found := s.handlers[kind]; found {
		return v
	}
	log.Panicf("Already registered Handler %q.", kind)

	return nil
}
