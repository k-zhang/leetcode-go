package easy

import "testing"

func Test_isSubSequence(t *testing.T) {
	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"abc", "ahbgdc"}, true},
		{"", args{"axc", "ahbgdc"}, false},
		{"", args{"", ""}, true},
		{"", args{"", "a"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubSequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
