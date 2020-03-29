package traverse

import (
	"reflect"
	"testing"

	"github.com/RossMerr/jsonschema"
)

func TestWalk(t *testing.T) {
	type args struct {
		s    *jsonschema.Schema
		path jsonschema.Path
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
				path: jsonschema.Path{},
			},
			want: &jsonschema.Schema{},
		},
		{
			name: "Definitions",
			args: args{
				s: &jsonschema.Schema{
					Definitions: map[string]*jsonschema.Schema{
						"test": &jsonschema.Schema{ID: jsonschema.ID("test")},
					},
				},
				path: jsonschema.Path{"definitions", "test"},
			},
			want: &jsonschema.Schema{ID: jsonschema.ID("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Walk(tt.args.s, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Walk() = %v, want %v", got, tt.want)
			}
		})
	}
}