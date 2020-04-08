package tags

import "testing"


func TestKeysString(t *testing.T) {
	type args struct {
		m map[string]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Key List",
			args: args{m: map[string]string{"foo": "bar", "hello": "world"}},
			want: []string{"foo=bar,hello=world", "hello=world,foo=bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KeysString(tt.args.m)
			if !(got == tt.want[0] || got == tt.want[1]) {
				t.Errorf("KeysString() = %v, want %v", got, tt.want)
			}
		})
	}
}
