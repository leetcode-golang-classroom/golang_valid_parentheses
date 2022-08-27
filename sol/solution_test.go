package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	s := "()[]{}"
	for idx := 0; idx < b.N; idx++ {
		isValid(s)
	}
}
func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "s = \"()\"",
			args: args{s: "()"},
			want: true,
		},
		{
			name: "s = \"()[]{}\"",
			args: args{s: "()[]{}"},
			want: true,
		},
		{
			name: "s = \"(]\"",
			args: args{s: "(]"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
