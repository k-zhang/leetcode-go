package easy

import "testing"

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"egg", "add"}, true},
		{"", args{"foo", "bar"}, false},
		{"", args{"paper", "title"}, true},
		{"", args{"c", "a"}, true},
		{"", args{"badc", "baba"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
