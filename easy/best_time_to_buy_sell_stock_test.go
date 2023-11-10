package easy

import "testing"

func Test_bestTimeToBuySellStock(t *testing.T) {
	type args struct {
		prices []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"", args{[]int{7, 6, 4, 3, 1}}, 0},
		{"", args{[]int{1}}, 0},
		{"", args{[]int{1, 1}}, 0},
		{"", args{[]int{1, 0, 1}}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestTimeToBuySellStock(tt.args.prices); got != tt.want {
				t.Errorf("bestTimeToBuySellStock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bestTimeToBuySellStock1(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
		{"", args{[]int{7, 6, 4, 3, 1}}, 0},
		{"", args{[]int{1}}, 0},
		{"", args{[]int{1, 1}}, 0},
		{"", args{[]int{1, 0, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestTimeToBuySellStock1(tt.args.prices); got != tt.want {
				t.Errorf("bestTimeToBuySellStock1() = %v, want %v", got, tt.want)
			}
		})
	}
}
