package interpreter

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

type Interpreter struct {
	config    *jsonschema.Config
	validator jsonschema.Validator
	parser    parser.Parser
}

func NewInterpreter(config *jsonschema.Config, validator jsonschema.Validator, parser parser.Parser) *Interpreter {
	return &Interpreter{
		config:    config,
		validator: validator,
		parser:    parser,
	}
}

func NewInterpreterDefaults(config *jsonschema.Config) *Interpreter {
	return NewInterpreter(config, jsonschema.NewValidator(), parser.NewParser(context.Background(), config.Packagename))
}

func (s *Interpreter) Interpret(files []string) (Interpret, error) {
	schemas := map[jsonschema.ID]*jsonschema.Schema{}

	for _, filename := range files {

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		var schema jsonschema.Schema
		json.Unmarshal(data, &schema)

		err = s.validator.ValidateSchema(s.config.Schemaversion, schema)
		if err != nil {
			return nil, err
		}

		schemas[schema.ID] = &schema
	}

	root := s.parser.Parse(schemas)

	return NewInterpretDefaults(root)
}
