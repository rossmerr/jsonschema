package jsonschema

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/RossMerr/jsonschema/functions"
	"github.com/RossMerr/jsonschema/types"
)

type SchemaReferences struct {
	Documents map[string]types.Document
	Config    Config
}

func NewSchemaReferences(config Config) *SchemaReferences {
	return &SchemaReferences{
		Config:    config,
		Documents: map[string]types.Document{},
	}
}

func (s *SchemaReferences) Parse(files []string) error {
	for _, filename := range files {

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		var doc types.Document
		json.Unmarshal(data, &doc)

		err = doc.ValidateSchema(s.Config.Schemaversion)
		if err != nil {
			return err
		}

		s.Documents[doc.ID] = doc
	}

	return nil
}

func (s SchemaReferences) Generate() error {

	templateBox, err := rice.FindBox("templates/")
	if err != nil {
		return err
	}

	schema, err := templateBox.String("schema.tmpl")
	if err != nil {
		return err
	}

	document, err := templateBox.String("document.tmpl")
	if err != nil {
		return err
	}

	object, err := templateBox.String("object.tmpl")
	if err != nil {
		return err
	}

	array, err := templateBox.String("array.tmpl")
	if err != nil {
		return err
	}

	string, err := templateBox.String("string.tmpl")
	if err != nil {
		return err
	}

	number, err := templateBox.String("number.tmpl")
	if err != nil {
		return err
	}

	validate, err := templateBox.String("validate.tmpl")
	if err != nil {
		return err
	}

	properties, err := templateBox.String("properties.tmpl")
	if err != nil {
		return err
	}

	reference, err := templateBox.String("reference.tmpl")
	if err != nil {
		return err
	}

	interfaces, err := templateBox.String("interface.tmpl")
	if err != nil {
		return err
	}

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
		}).Parse(schema)

		tmpl, err = tmpl.Parse(document)
		tmpl, err = tmpl.Parse(object)
		tmpl, err = tmpl.Parse(array)
		tmpl, err = tmpl.Parse(string)
		tmpl, err = tmpl.Parse(number)
		tmpl, err = tmpl.Parse(validate)
		tmpl, err = tmpl.Parse(properties)
		tmpl, err = tmpl.Parse(reference)
		tmpl, err = tmpl.Parse(interfaces)

		if err != nil {
			return err
		}

		filename := s.Config.Output + doc.Filename() + ".go"
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

		schema := Schema{
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
