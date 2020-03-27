package traversal

import (
	"reflect"
	"testing"

	"github.com/RossMerr/jsonschema"
)

func TestTraverse(t *testing.T) {
	type args struct {
		s         *jsonschema.Schema
		reference jsonschema.Reference
	}
	tests := []struct {
		name string
		args args
		want *jsonschema.Schema
	}{
		{
			name: "Empty",
			args: args{
				s:         &jsonschema.Schema{},
				reference: jsonschema.Reference(""),
			},
			want: nil,
		},
		{
			name: "Definitions",
			args: args{
				s: &jsonschema.Schema{
					Definitions: map[string]*jsonschema.Schema{
						"test": &jsonschema.Schema{ID: jsonschema.ID("test")},
					},
				},
				reference: jsonschema.Reference("#/Definitions/test"),
			},
			want: &jsonschema.Schema{ID: jsonschema.ID("test")},
		},
		{
			name: "Pointer",
			args: args{
				s: &jsonschema.Schema{
					OneOf: []*jsonschema.Schema{
						&jsonschema.Schema{ID: jsonschema.ID("test")},
					},
				},
				reference: jsonschema.Reference("#test"),
			},
			want: &jsonschema.Schema{ID: jsonschema.ID("test")},
		},
	}
	{
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Traverse(tt.args.s, tt.args.reference); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Traverse() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
