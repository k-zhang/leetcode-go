package easy

import (
	"reflect"
	"testing"
)

func Test_plus_one(t *testing.T) {
	type args struct {
		digits []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{0}}, []int{1}},
		{"", args{[]int{1, 2, 3}}, []int{1, 2, 4}},
		{"", args{[]int{4, 3, 2, 1}}, []int{4, 3, 2, 2}},
		{"", args{[]int{4, 3, 2, 9}}, []int{4, 3, 3, 0}},
		{"", args{[]int{4, 3, 9, 9}}, []int{4, 4, 0, 0}},
		{"", args{[]int{4, 9, 9, 9}}, []int{5, 0, 0, 0}},
		{"", args{[]int{9, 9, 9, 9}}, []int{1, 0, 0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plus_one() = %v, want %v", got, tt.want)
			}
		})
	}
}
