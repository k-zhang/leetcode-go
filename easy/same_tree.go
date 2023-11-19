package easy

import "github.com/k-zhang/leetcode-go/common"

func isSameTree(p, q *common.TreeNode[int]) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
