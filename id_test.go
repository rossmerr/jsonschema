package jsonschema

import (
	"testing"
)

func TestID_Filename(t *testing.T) {
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
			want: "test",
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
			want: "test",
		},
		{
			name: "Absolute",
			s:    ID("http://www.sample.com/test.json#defintions/hello"),
			want: "test",
		},
		{
			name: "Test Case",
			s:    ID("test.json#defintions/hello_world"),
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToFilename(); got != tt.want {
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
			s: ID(""),
			want:".",
		},
		{
			name: "Fragment",
			s: ID("#test"),
			want:".",
		},
		{
			name: "Relative",
			s: ID("/test/#test"),
			want:".",
		},
		{
			name: "Absolute",
			s: ID("http://www.sample.com"),
			want:"http://www.sample.com",
		},
		{
			name: "Absolute with path",
			s: ID("http://www.sample.com/foo/bar/"),
			want:"http://www.sample.com/foo/bar",
		},
		{
			name: "Absolute with path and fragment",
			s: ID("http://www.sample.com/foo#test"),
			want:"http://www.sample.com/foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanonicalURI(tt.s.String()); got != tt.want {
				t.Errorf("CanonicalURI() = %v, want %v", got, tt.want)
			}
		})
	}
}