package checker

import "testing"

func TestCheckArg(t *testing.T) {
	type args struct {
		arg interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test int",
			args: args{arg: 1},
			want: "1:type is int.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckArg(tt.args.arg); got != tt.want {
				t.Errorf("CheckArg() = %v, want %v", got, tt.want)
			}
		})
	}
}
