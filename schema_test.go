package jsonschema_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/types"
)

func TestSchemas_Generate(t *testing.T) {
	type fields struct {
		documents map[string]types.Document
		config    jsonschema.Config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		// {
		// 	name: "Basic",
		// 	fields: fields{
		// 		documents: map[string]types.Document{
		// 			"id": loadRawSchema("samples/basicBasic.json"),
		// 		},
		// 		config: jsonschema.Config{
		// 			Packagename: "main",
		// 			Output: "output/",
		// 		},
		// 	},
		// },
		// {
		// 	name: "Nesting data structures",
		// 	fields: fields{
		// 		documents: map[string]types.Document{
		// 			"id": loadRawSchema("samples/productNesting.json"),
		// 		},
		// 		config: jsonschema.Config{
		// 			Packagename: "main",
		// 			Output: "output/",
		// 		},
		// 	},
		// },
		// {
		// 	name: "References outside the schema",
		// 	fields: fields{
		// 		documents: map[string]types.Document{
		// 			"https://example.com/geographical-location.schema.json": loadRawSchema("samples/geographical-location.schema.json"),
		// 			"http://example.com/product.schema.json": loadRawSchema("samples/product.schema.json"),
		//
		// 		},
		// 		config: jsonschema.Config{
		// 			Packagename: "main",
		// 			Output: "output/",
		// 		},
		// 	},
		// },
		{
			name: "Oneof",
			fields: fields{
				documents: map[string]types.Document{
					"http://example.com/entry-schema": loadRawSchema("samples/entry-schema.json"),
				},
				config: jsonschema.Config{
					Packagename: "main",
					Output:      "output/",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &jsonschema.SchemaReferences{
				Documents: tt.fields.documents,
				Config:    tt.fields.config,
			}
			if err := s.Generate(); (err != nil) != tt.wantErr {
				t.Errorf("Schemas.Generate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func loadRawSchema(filename string) types.Document {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var doc types.Document
	err = json.Unmarshal(data, &doc)
	if err != nil {
		panic(err)
	}

	return doc
}
