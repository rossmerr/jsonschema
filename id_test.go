package jsonschema_test

import (
	"testing"

	"github.com/RossMerr/jsonschema"
)

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
