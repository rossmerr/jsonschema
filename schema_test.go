package jsonschema_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/interpreter"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/handlers"
)

func TestSchemas_Generate(t *testing.T) {

	os.MkdirAll("output/", 0755)

	type fields struct {
		documents map[jsonschema.ID]*jsonschema.Schema
		paths     []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Basic",
			fields: fields{
				paths: []string{"samples/basic.json"},
			},
		},
		{
			name: "Enum",
			fields: fields{
				paths: []string{"samples/enum.json"},
			},
		},
		{
			name: "Nesting data structures",
			fields: fields{
				paths: []string{"samples/nesting.json"},
			},
		},
		{
			name: "References inside the schema",
			fields: fields{
				paths: []string{"samples/reference.json"},
			},
		},
		{
			name: "References outside the schema",
			fields: fields{
				paths: []string{"samples/reference-outside.schema.json", "samples/reference-outside.json"},
			},
		},
		{
			name: "Oneof",
			fields: fields{
				paths: []string{"samples/oneof.json"},
			},
		},
		{
			name: "AnyOf",
			fields: fields{
				paths: []string{"samples/anyof.json"},
			},
		},
		{
			name: "AllOf",
			fields: fields{
				paths: []string{"samples/allof.json"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			files := []string{}
			p := parser.NewParser("main")
			p = handlers.DefaultHandlers(p)
			documents := map[jsonschema.ID]jsonschema.JsonSchema{}
			references := map[jsonschema.ID]jsonschema.JsonSchema{}
			for _, path := range tt.fields.paths {
				data, err := ioutil.ReadFile(path)
				if err != nil {
					panic(err)
				}

				var schema jsonschema.RootSchema
				err = json.Unmarshal(data, &schema)
				if err != nil {
					panic(err)
				}
				refs := jsonschema.ResolveIDs(data)

				documents[schema.ID] = &schema
				for k, v := range refs {
					references[k] = v
				}
			}

			parse, err := p.Parse(documents, references)
			if err != nil {
				t.Error(err)
			}

			interpret, err := interpreter.NewInterpretDefaults(parse)
			if err != nil {
				t.Error(err)
			}
			if files, err = interpret.ToFile("output/"); (err != nil) != tt.wantErr {
				t.Errorf("Schemas.Generate() error = %v, wantErr %v", err, tt.wantErr)
				files = []string{}
			}

			t.Cleanup(func() {
				for _, file := range files {
					fmt.Printf("%v", file)
					// if err := os.Remove(file); err != nil {
					// 	t.Error("error resetting:", err)
					// }
				}
			})
		})
	}

	// t.Cleanup(func() {
	// 	if !t.Failed() {
	// 		os.RemoveAll("output/")
	// 	}
	// })
}
