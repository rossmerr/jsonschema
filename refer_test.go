package jsonschema_test

import (
	"testing"

	"github.com/RossMerr/jsonschema"
)

func TestRefer_Typename(t *testing.T) {
	tests := []struct {
		name string
		s    jsonschema.Refer
		want string
	}{
		{
			name:"Definitions",
			s: jsonschema.Refer("#/definitions/diskDevice"),
			want: "DiskDevice",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Typename(); got != tt.want {
				t.Errorf("Typename() = %v, want %v", got, tt.want)
			}
		})
	}
}