package interpreter

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/gookit/color"

	log "github.com/sirupsen/logrus"
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
	references := map[jsonschema.ID]*jsonschema.Schema{}
	green := color.FgCyan.Render
	red := color.FgRed.Render

	rawFiles := map[string][]byte{}
	for _, filename := range files {

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf(red("🗴") + "Reading file\n")
			return nil, err
		}

		log.Infof("Found file %v", filename)
		rawFiles[filename] = data
	}

	fmt.Printf(green("✓")+" Found %v files\n", len(rawFiles))

	i := 0
	for _, data := range rawFiles {
		var schema jsonschema.Schema
		err := json.Unmarshal(data, &schema)
		if err != nil {
			fmt.Printf(red("🗴") + "Unmarshalling\n")
			return nil, err
		} else {
			i++
		}
		log.Infof("Found schema %v", schema.ID)

		schemas[schema.ID] = &schema
	}

	fmt.Printf(green("✓")+" Unmarshalled %v schemas\n", i)

	for _, data := range rawFiles {

		refs := jsonschema.ResolveIDs(data)

		for k, v := range refs {
			references[k] = v
			log.Infof("Found reference %v", k)
		}
	}

	fmt.Printf(green("✓")+" Found %v references\n", len(references))

	for id, schema := range schemas {
		err := s.validator.ValidateSchema(id, schema)
		if err != nil {
			fmt.Printf(red("🗴") + "Validating schemas\n")
			return nil, err
		}
	}

	fmt.Printf(green("✓") + " Schemas valid \n")

	root := s.parser.Parse(schemas, references)

	return NewInterpretDefaults(root)
}
