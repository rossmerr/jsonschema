package jsonschema_test

import (
	"testing"

	"github.com/RossMerr/jsonschema"
)

func TestMetaSchemasToString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Draft 2019-09",
			want: "https://json-schema.org/2019-09/schema",
		},
		{
			name: "Draft 08 (Old name)",
			want: "http://json-schema.org/draft-08/schema",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !jsonschema.Contains(jsonschema.MetaSchemasToString(), tt.want) {
				t.Errorf("MetaSchemasToString() =  want %v", tt.want)
			}
		})
	}
}
