package interpreter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/handlers"
	log "github.com/sirupsen/logrus"
)

type Interpreter struct {
	parser parser.Parser
}

func NewInterpreter(parser parser.Parser) *Interpreter {
	return &Interpreter{
		parser: parser,
	}
}

func NewInterpreterDefaults(packagename string) *Interpreter {
	p := parser.NewParser(packagename)
	p = handlers.DefaultHandlers(p)
	return NewInterpreter(p)
}

func (s *Interpreter) Interpret(files []string) (Interpret, error) {
	references := map[jsonschema.ID]*jsonschema.Schema{}

	rawFiles, err := s.readFiles(files)
	if err != nil {
		return nil, err
	}

	schemas, err := s.unmarshall(rawFiles)
	if err != nil {
		return nil, err
	}

	err = s.references(rawFiles, references)
	if err != nil {
		return nil, err
	}

	root, err := s.parser.Parse(schemas, references)
	if err != nil {
		return nil, err
	}
	return NewInterpretDefaults(root)
}

func (s *Interpreter) references(rawFiles map[string][]byte, references map[jsonschema.ID]*jsonschema.Schema) (error) {
	for _, data := range rawFiles {

		refs := jsonschema.ResolveIDs(data)

		for k, v := range refs {
			if _, ok := references[k]; !ok {
				references[k] = v
			} else {
				fmt.Printf(red("🗴") + "References\n")

				return fmt.Errorf("interpreter: Reference keys need to be unique found %v more than once", k)
			}
			log.Infof("Found reference %v", k)
		}
	}

	fmt.Printf(green("✓")+" Found %v references\n", len(references))
	return nil
}

func (s *Interpreter) unmarshall(rawFiles map[string][]byte) ( map[jsonschema.ID]*jsonschema.Schema, error) {
	schemas := map[jsonschema.ID]*jsonschema.Schema{}
	i := 0
	for _, data := range rawFiles {
		var schema jsonschema.Schema
		err := json.Unmarshal(data, &schema)
		if err != nil {
			fmt.Printf(red("🗴") + "Unmarshalling\n")
			return schemas, err
		} else {
			i++
		}
		log.Infof("Found schema %v", schema.ID)

		schemas[schema.ID] = &schema
	}

	fmt.Printf(green("✓")+" Unmarshalled %v schemas\n", i)
	return schemas, nil
}

func (s *Interpreter) readFiles(files []string) (map[string][]byte, error) {
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
	return rawFiles, nil
}
