package common

type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func NewTreeNode[T any](val T) *TreeNode[T] {
	return &TreeNode[T]{
		val,
		nil,
		nil,
	}
}

func (t *TreeNode[T]) AddLeft(val T) *TreeNode[T] {
	left := TreeNode[T]{
		val,
		nil,
		nil,
	}

	t.Left = &left

	return t
}

func (t *TreeNode[T]) AddRight(val T) *TreeNode[T] {
	right := TreeNode[T]{
		val,
		nil,
		nil,
	}

	t.Right = &right

	return t
}
