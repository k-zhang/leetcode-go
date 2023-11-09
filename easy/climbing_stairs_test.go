package easy

import "testing"

func Test_climbingStairs(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{1}, 1},
		{"", args{2}, 2},
		{"", args{3}, 3},
		{"", args{4}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbingStairs(tt.args.n); got != tt.want {
				t.Errorf("climbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
