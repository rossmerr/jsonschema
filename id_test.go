package jsonschema_test

import (
	"testing"

	"github.com/RossMerr/jsonschema"
)

func TestID_ToTypename(t *testing.T) {
	tests := []struct {
		name string
		s    jsonschema.ID
		want string
	}{
		{
			name: "Empty string",
			s:    jsonschema.ID(""),
			want: ".",
		},
		{
			name: "No #",
			s:    jsonschema.ID("test"),
			want: "Test",
		},
		{
			name: "Root ID",
			s:    jsonschema.ID("#test"),
			want: ".",
		},
		{
			name: "Defintions",
			s:    jsonschema.ID("#defintions/test"),
			want: ".",
		},
		{
			name: "Relative",
			s:    jsonschema.ID("test.json#defintions/hello"),
			want: "Test",
		},
		{
			name: "Absolute",
			s:    jsonschema.ID("http://www.sample.com/test.json#defintions/hello"),
			want: "Test",
		},
		{
			name: "Test Case",
			s:    jsonschema.ID("test.json#defintions/hello_world"),
			want: "Test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToTypename(); got != tt.want {
				t.Errorf("ToTypename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestID_CanonicalURI(t *testing.T) {
	type args struct {
		id  jsonschema.ID
		err error
	}

	newArgs := func(id jsonschema.ID, err error) args {
		return args{
			id:  id,
			err: err,
		}
	}
	tests := []struct {
		name string
		s    args
		want string
	}{
		{
			name: "Empty",
			s:    newArgs(jsonschema.NewID("")),
			want: ".",
		},
		{
			name: "Fragment",
			s:    newArgs(jsonschema.NewID("#test")),
			want: ".",
		},
		{
			name: "Relative",
			s:    newArgs(jsonschema.NewID("/test/#test")),
			want: ".",
		},
		{
			name: "Absolute",
			s:    newArgs(jsonschema.NewID("http://www.sample.com")),
			want: "http://www.sample.com",
		},
		{
			name: "Absolute with path",
			s:    newArgs(jsonschema.NewID("http://www.sample.com/foo/bar/")),
			want: "http://www.sample.com/foo/bar",
		},
		{
			name: "Absolute with path and fragment",
			s:    newArgs(jsonschema.NewID("http://www.sample.com/foo#test")),
			want: "http://www.sample.com/foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.id.String(); got != tt.want {
				t.Errorf("CanonicalURI() = %v, want %v", got, tt.want)
			}
		})
	}
}
