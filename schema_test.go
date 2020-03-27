package jsonschema_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/interpreter"
	"github.com/RossMerr/jsonschema/parser"
)

func TestSchemas_Generate(t *testing.T) {
	type fields struct {
		documents map[jsonschema.ID]*jsonschema.Schema
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{
			name: "Basic",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"basicBasic": loadRawSchema("samples/basic.json"),
				},
			},
		},
		{
			name: "Nesting data structures",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"productNesting": loadRawSchema("samples/nesting.json"),
				},
			},
		},
		{
			name: "References inside the schema",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"http://example.com/reference.json": loadRawSchema("samples/reference.json"),
				},
			},
		},
		{
			name: "References outside the schema",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"https://example.com/reference-outside.schema.json": loadRawSchema("samples/reference-outside.schema.json"),
					"http://example.com/reference-outside.json":         loadRawSchema("samples/reference-outside.json"),
				},
			},
		},
		{
			name: "Oneof",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"http://example.com/oneof.json": loadRawSchema("samples/oneof.json"),
				},
			},
		},
		{
			name: "AnyOf",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"http://example.com/anyof.json": loadRawSchema("samples/anyof.json"),
				},
			},
		},
		{
			name: "AllOf",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"http://example.com/allof.json": loadRawSchema("samples/allof.json"),
				},
			},
		},
	}
	for _, tt := range tests {
		os.MkdirAll("output/", 0755)
		t.Run(tt.name, func(t *testing.T) {

			p := parser.NewParser(context.Background(), "main")
			parse := p.Parse(tt.fields.documents)
			interpret, err := interpreter.NewInterpretDefaults(parse)
			if err != nil {
				t.Error(err)
			}
			if err := interpret.ToFile("output/"); (err != nil) != tt.wantErr {
				t.Errorf("Schemas.Generate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	t.Cleanup(func() {
		if err := os.RemoveAll("output/"); err != nil {
			t.Error("error resetting:", err)
		}
	})
}

func loadRawSchema(filename string) *jsonschema.Schema {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var doc jsonschema.Schema
	err = json.Unmarshal(data, &doc)
	if err != nil {
		panic(err)
	}

	return &doc
}
