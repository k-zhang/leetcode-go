package easy

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{0}, true},
		{"", args{5}, true},
		{"", args{121}, true},
		{"", args{-121}, false},
		{"", args{10}, false},
		{"", args{2147483647}, false},
		{"", args{-2147483648}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
