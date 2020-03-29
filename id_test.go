package jsonschema

import (
	"testing"
)

func TestID_ToTypename(t *testing.T) {
	tests := []struct {
		name string
		s    ID
		want string
	}{
		{
			name: "Empty string",
			s:    ID(""),
			want: ".",
		},
		{
			name: "No #",
			s:    ID("test"),
			want: "Test",
		},
		{
			name: "Root ID",
			s:    ID("#test"),
			want: ".",
		},
		{
			name: "Defintions",
			s:    ID("#defintions/test"),
			want: ".",
		},
		{
			name: "Relative",
			s:    ID("test.json#defintions/hello"),
			want: "Test",
		},
		{
			name: "Absolute",
			s:    ID("http://www.sample.com/test.json#defintions/hello"),
			want: "Test",
		},
		{
			name: "Test Case",
			s:    ID("test.json#defintions/hello_world"),
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
	tests := []struct {
		name string
		s    ID
		want string
	}{
		{
			name: "Empty",
			s:    NewID(""),
			want: ".",
		},
		{
			name: "Fragment",
			s:    NewID("#test"),
			want: ".",
		},
		{
			name: "Relative",
			s:    NewID("/test/#test"),
			want: ".",
		},
		{
			name: "Absolute",
			s:    NewID("http://www.sample.com"),
			want: "http://www.sample.com",
		},
		{
			name: "Absolute with path",
			s:    NewID("http://www.sample.com/foo/bar/"),
			want: "http://www.sample.com/foo/bar",
		},
		{
			name: "Absolute with path and fragment",
			s:    NewID("http://www.sample.com/foo#test"),
			want: "http://www.sample.com/foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("CanonicalURI() = %v, want %v", got, tt.want)
			}
		})
	}
}
