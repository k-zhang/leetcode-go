package easy

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"abbaca"}, "ca"},
		{"", args{"azxxzy"}, "ay"},
		{"", args{"aa"}, ""},
		{"", args{"bcaacb"}, ""},
		{"", args{"aac"}, "c"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.s); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
