package jsonschema_test

import (
	"reflect"
	"testing"

	"github.com/RossMerr/jsonschema"
)

func TestReference_Fragments(t *testing.T) {
	type args struct {
		ref jsonschema.Reference
		err error
	}

	newArgs := func(ref jsonschema.Reference, err error) args {
		return args{
			ref: ref,
			err: err,
		}
	}

	tests := []struct {
		name      string
		s         args
		wantQuery jsonschema.Path
	}{
		{
			name:      "Empty string",
			s:         newArgs(jsonschema.NewReference("")),
			wantQuery: jsonschema.Path{},
		},
		{
			name:      "No pointer",
			s:         newArgs(jsonschema.NewReference("test")),
			wantQuery: jsonschema.Path{},
		},
		{
			name:      "Just pointer",
			s:         newArgs(jsonschema.NewReference("#")),
			wantQuery: jsonschema.Path{},
		},
		{
			name:      "No path in fragment",
			s:         newArgs(jsonschema.NewReference("#test")),
			wantQuery: jsonschema.Path{},
		},
		{
			name:      "ID followed by fragments",
			s:         newArgs(jsonschema.NewReference("#test/hello/world")),
			wantQuery: jsonschema.Path{"hello", "world"},
		},
		{
			name:      "Two fragment",
			s:         newArgs(jsonschema.NewReference("#/hello/world")),
			wantQuery: jsonschema.Path{"hello", "world"},
		},
		{
			name:      "Relative",
			s:         newArgs(jsonschema.NewReference("test.json#/hello/world")),
			wantQuery: jsonschema.Path{"hello", "world"},
		},
		{
			name:      "Relative",
			s:         newArgs(jsonschema.NewReference("http://www.sample.com/test.json#/hello/world")),
			wantQuery: jsonschema.Path{"hello", "world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if path := tt.s.ref.Path(); !reflect.DeepEqual(path, tt.wantQuery) {
				t.Errorf("Path() = %v, want %v", path, tt.wantQuery)
			}
		})
	}
}

func TestReference_Pointer(t *testing.T) {
	type args struct {
		ref jsonschema.Reference
		err error
	}

	newArgs := func(ref jsonschema.Reference, err error) args {
		return args{
			ref: ref,
			err: err,
		}
	}

	tests := []struct {
		name string
		s    args
		want jsonschema.ID
	}{
		{
			name: "Empty string",
			s:    newArgs(jsonschema.NewReference("")),
			want: jsonschema.ID("."),
		},
		{
			name: "No pointer",
			s:    newArgs(jsonschema.NewReference("test")),
			want: jsonschema.ID("."),
		},
		{
			name: "Just pointer",
			s:    newArgs(jsonschema.NewReference("#")),
			want: jsonschema.ID("."),
		},
		{
			name: "ID",
			s:    newArgs(jsonschema.NewReference("#test")),
			want: jsonschema.ID("."),
		},
		{
			name: "One fragment no pointer",
			s:    newArgs(jsonschema.NewReference("#/test")),
			want: jsonschema.ID("."),
		},
		{
			name: "ID followed by two fragment",
			s:    newArgs(jsonschema.NewReference("#test/hello/world")),
			want: jsonschema.ID("."),
		},
		{
			name: "ID followed by two fragment",
			s:    newArgs(jsonschema.NewReference("https://www.sample.com/#test/hello/world")),
			want: jsonschema.ID("https://www.sample.com"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := tt.s.ref.ID(); got != tt.want {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}
