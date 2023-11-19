package easy

import (
	"github.com/k-zhang/leetcode-go/common"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *common.TreeNode[int]
		q *common.TreeNode[int]
	}

	array1 := common.NewArrayWithNullableElement([]int{1, 2, 3}, []int{})
	array2 := common.NewArrayWithNullableElement([]int{1, 2, 3}, []int{})
	array3 := common.NewArrayWithNullableElement([]int{1, 2}, []int{})
	array4 := common.NewArrayWithNullableElement([]int{1, -1, 2}, []int{1})
	array5 := common.NewArrayWithNullableElement([]int{1, 2, 1}, []int{})
	array6 := common.NewArrayWithNullableElement([]int{1, 1, 2}, []int{})
	array7 := common.NewArrayWithNullableElement([]int{}, []int{})
	array8 := common.NewArrayWithNullableElement([]int{}, []int{})
	array9 := common.NewArrayWithNullableElement([]int{1}, []int{})
	array10 := common.NewArrayWithNullableElement([]int{1}, []int{})
	array11 := common.NewArrayWithNullableElement([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{})
	array12 := common.NewArrayWithNullableElement([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{})

	tree1 := common.ConstructTree[int](&array1)
	tree2 := common.ConstructTree[int](&array2)
	tree3 := common.ConstructTree[int](&array3)
	tree4 := common.ConstructTree[int](&array4)
	tree5 := common.ConstructTree[int](&array5)
	tree6 := common.ConstructTree[int](&array6)
	tree7 := common.ConstructTree[int](&array7)
	tree8 := common.ConstructTree[int](&array8)
	tree9 := common.ConstructTree[int](&array9)
	tree10 := common.ConstructTree[int](&array10)
	tree11 := common.ConstructTree[int](&array11)
	tree12 := common.ConstructTree[int](&array12)

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{tree1, tree2}, true},
		{"", args{tree3, tree4}, false},
		{"", args{tree5, tree6}, false},
		{"", args{tree7, tree8}, true},
		{"", args{tree9, tree10}, true},
		{"", args{tree11, tree12}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
