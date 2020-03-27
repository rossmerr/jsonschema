package interpreter

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

type Interpreter struct {
	validator jsonschema.Validator
	parser    parser.Parser
}

func NewInterpreter(validator jsonschema.Validator, parser parser.Parser) *Interpreter {
	return &Interpreter{
		validator: validator,
		parser:    parser,
	}
}

func NewInterpreterDefaults(packagename string) *Interpreter {
	return NewInterpreter(jsonschema.NewValidator(), parser.NewParser(context.Background(), packagename))
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

		err = s.validator.ValidateSchema(schema)
		if err != nil {
			return nil, err
		}

		schemas[schema.ID] = &schema
	}

	root := s.parser.Parse(schemas)

	return NewInterpretDefaults(root)
}
