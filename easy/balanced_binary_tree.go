package easy

import (
	"github.com/k-zhang/leetcode-go/common"
	"math"
)

func isBalanced(root *common.TreeNode[int]) bool {
	if root == nil {
		return true
	}

	return dfs(root, 0) != -1
}

func dfs(root *common.TreeNode[int], depth int) int {
	if root == nil {
		return depth
	}

	depthLeft := dfs(root.Left, depth+1)
	depthRight := dfs(root.Right, depth+1)

	if depthLeft == -1 || depthRight == -1 {
		return -1
	}

	if math.Abs(float64(depthLeft-depthRight)) <= 1 {
		return int(math.Max(float64(depthLeft), float64(depthRight)))
	} else {
		return -1
	}
}
