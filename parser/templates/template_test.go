package templates

import "testing"


func TestTypename(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty string",
			args: args{s: ""},
			want: "",
		},
		{
			name: "Unicode letter",
			args: args{s: "foo bar"},
			want: "FooBar",
		},
		{
			name: "Starting with a none unicode letter",
			args: args{s: "1 foo bar"},
			want: "No1FooBar",
		},
		{
			name: "Testing dash",
			args: args{s: "foo-bar"},
			want: "FooBar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Typename(tt.args.s); got != tt.want {
				t.Errorf("Structname() = %v, want %v", got, tt.want)
			}
		})
	}
}
