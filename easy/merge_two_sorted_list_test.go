package easy

import (
	"github.com/k-zhang/leetcode-go/common"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *common.ListNode[int]
		l2 *common.ListNode[int]
	}

	left1 := common.ArrayToLinkedList[int]([]int{1, 2, 4})
	right1 := common.ArrayToLinkedList[int]([]int{1, 3, 4})
	expect1 := common.ArrayToLinkedList[int]([]int{1, 1, 2, 3, 4, 4})

	left2 := common.ArrayToLinkedList[int]([]int{})
	right2 := common.ArrayToLinkedList[int]([]int{})
	expect2 := common.ArrayToLinkedList[int]([]int{})

	left3 := common.ArrayToLinkedList[int]([]int{})
	right3 := common.ArrayToLinkedList[int]([]int{1})
	expect3 := common.ArrayToLinkedList[int]([]int{1})

	left4 := common.ArrayToLinkedList[int]([]int{1})
	right4 := common.ArrayToLinkedList[int]([]int{})
	expect4 := common.ArrayToLinkedList[int]([]int{1})

	left5 := common.ArrayToLinkedList[int]([]int{1, 1, 1, 1})
	right5 := common.ArrayToLinkedList[int]([]int{1, 1, 1, 1})
	expect5 := common.ArrayToLinkedList[int]([]int{1, 1, 1, 1, 1, 1, 1, 1})

	left6 := common.ArrayToLinkedList[int]([]int{2, 4, 6, 8, 10, 12, 14})
	right6 := common.ArrayToLinkedList[int]([]int{1, 3, 5, 7, 9, 11, 13, 15})
	expect6 := common.ArrayToLinkedList[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})

	left7 := common.ArrayToLinkedList[int]([]int{2, 4, 6, 8, 10, 12, 14})
	right7 := common.ArrayToLinkedList[int]([]int{9})
	expect7 := common.ArrayToLinkedList[int]([]int{2, 4, 6, 8, 9, 10, 12, 14})

	tests := []struct {
		name string
		args args
		want *common.ListNode[int]
	}{
		{"", args{left1, right1}, expect1},
		{"", args{left2, right2}, expect2},
		{"", args{left3, right3}, expect3},
		{"", args{left4, right4}, expect4},
		{"", args{left5, right5}, expect5},
		{"", args{left6, right6}, expect6},
		{"", args{left7, right7}, expect7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
