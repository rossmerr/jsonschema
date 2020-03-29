package findReferences

import (
	"reflect"
	"testing"

	"github.com/RossMerr/jsonschema"
)

func TestWalk(t *testing.T) {
	type args struct {
		s *jsonschema.Schema
	}
	tests := []struct {
		name string
		args args
		want map[jsonschema.ID]*jsonschema.Schema
	}{
		{
			name: "Schema identification",
			args: args{s: &jsonschema.Schema{
				ID: "http://example.com/root.json",
				Defs: map[string]*jsonschema.Schema{
					"A": &jsonschema.Schema{
						ID: "https://example.com/foo",
					},
					"B": &jsonschema.Schema{
						ID: "other.json",
						Defs: map[string]*jsonschema.Schema{
							"X": &jsonschema.Schema{ID: "#bar"},
							"Y": &jsonschema.Schema{ID: "https://example.com/bar#test"},
						},
					},
				},
			},
			},
			want: map[jsonschema.ID]*jsonschema.Schema{
				jsonschema.ID("http://example.com/root.json"): &jsonschema.Schema{
					ID: "http://example.com/root.json",
					Defs: map[string]*jsonschema.Schema{
						"A": &jsonschema.Schema{
							ID: "https://example.com/foo",
						},
						"B": &jsonschema.Schema{
							ID: "other.json",
							Defs: map[string]*jsonschema.Schema{
								"X": &jsonschema.Schema{ID: "#bar"},
								"Y": &jsonschema.Schema{ID: "https://example.com/bar#test"},
							},
						},
					},
				},
				jsonschema.ID("https://example.com/bar"): &jsonschema.Schema{
					ID: "https://example.com/bar#test",
				},
				jsonschema.ID("https://example.com/foo"): &jsonschema.Schema{
					ID: "https://example.com/foo",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Walk(tt.args.s)
			for k, want := range tt.want {
				got := result[k]
				if !reflect.DeepEqual(got, want) {
					t.Errorf("Walk() = %v, want %v", got, want)
				}
			}
		})
	}
}
