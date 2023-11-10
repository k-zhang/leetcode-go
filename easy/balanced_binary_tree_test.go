package easy

import (
	"github.com/k-zhang/leetcode-go/common"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *common.TreeNode[int]
	}

	var input1 = common.NewArrayWithNullableElement[int]([]int{3, 9, 20, -1, -1, 15, 7}, []int{3, 4})
	var input2 = common.NewArrayWithNullableElement[int]([]int{3, 9, 20, -1, -1, -1, 7}, []int{3, 4, 5})
	var input3 = common.NewArrayWithNullableElement[int]([]int{3, 9, 20, -1, -1, -1, 7, 8, -1}, []int{3, 4, 5, 8})
	var input4 = common.NewArrayWithNullableElement[int]([]int{1, 2, 2, 3, 3, -1, -1, 4, 4}, []int{5, 6})
	var input5 = common.NewArrayWithNullableElement[int]([]int{}, []int{})

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{common.ConstructTree[int](&input1)}, true},
		{"", args{common.ConstructTree[int](&input2)}, true},
		{"", args{common.ConstructTree[int](&input3)}, false},
		{"", args{common.ConstructTree[int](&input4)}, false},
		{"", args{common.ConstructTree[int](&input5)}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
