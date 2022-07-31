package strings

import "testing"

func TestIsIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"is isomorphic",
			args{
				"egg",
				"add",
			},
			true,
		},
		{
			"isn't isomorphic",
			args{
				"foo",
				"bar",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"is subsequence",
			args{
				"abc",
				"ahbgdc",
			},
			true,
		},
		{
			"isn't subsequence",
			args{
				"axc",
				"ahbgdc",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
