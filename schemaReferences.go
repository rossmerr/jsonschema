package jsonschema

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"

	"github.com/RossMerr/jsonschema/functions"
)

type SchemaReferences struct {
	Documents map[string]*Schema
	Config    Config
}

func NewSchemaReferences(config Config) *SchemaReferences {
	return &SchemaReferences{
		Config:    config,
		Documents: map[string]*Schema{},
	}
}

func (s *SchemaReferences) Parse(files []string) error {
	for _, filename := range files {

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		var schema Schema
		json.Unmarshal(data, &schema)

		err = ValidateSchema(s.Config.Schemaversion, schema)
		if err != nil {
			return err
		}

		s.Documents[schema.ID] = &schema
	}


	return nil
}

func (s SchemaReferences) Generate() error {

	for id, doc := range s.Documents {
		tmpl, err := template.New("schema.tmpl").Funcs(template.FuncMap{
			"toString":      functions.ToString,
			"typename":      functions.Typename,
			"isStruct":      functions.IsStruct,
			"isArray":       functions.IsArray,
			"isString":      functions.IsString,
			"isNumber":      functions.IsNumber,
			"isInterface":   functions.IsInterface,
			"dict":          functions.Dict,
			"isRequired":    functions.IsRequired,
			"isNotRequired": functions.IsNotRequired,
			"validate":      functions.Validate,
			"add":           functions.Add,
			"isPointer":     functions.IsPointer,
			"titleCase":     functions.TitleCase,
			"mixedCase":     functions.MixedCase,
		}).ParseFiles(
			"templates/schema.tmpl",
			"templates/document.tmpl",
			"templates/object.tmpl",
			"templates/array.tmpl",
			"templates/string.tmpl",
			"templates/number.tmpl",
			"templates/validate.tmpl",
			"templates/properties.tmpl",
			"templates/reference.tmpl",
			"templates/interface.tmpl",
		)
		if err != nil {
			return err
		}

		filename := s.Config.Output + functions.Filename(doc.ID) + ".go"
		_, err = os.Stat(filename)
		if !os.IsNotExist(err) {
			err = os.Remove(filename)
			if err != nil {
				return err
			}
		}

		file, err := os.Create(filename)
		if err != nil {
			return err
		}

		schema := Mapping{
			ID:          id,
			Document:    doc,
			Config:      s.Config,
			Definitions: &s,
		}
		err = tmpl.Execute(file, schema)
		if err != nil {
			return err
		}

		cmd := exec.Command("gofmt", "-w", filename)
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
