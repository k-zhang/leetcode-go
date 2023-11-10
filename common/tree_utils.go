package common

type ArrayWithNullableElement[T any] struct {
	size   int
	values map[int]T
}

func (array *ArrayWithNullableElement[T]) Add(index int, value T) {
	array.size = array.size + 1
	array.values[index] = value
}

func (array *ArrayWithNullableElement[T]) Get(index int) (T, bool) {
	val, prs := array.values[index]
	return val, prs
}

func NewArrayWithNullableElement[T any](input []T, missingValueIndices []int) ArrayWithNullableElement[T] {
	var array = ArrayWithNullableElement[T]{
		len(input),
		make(map[int]T),
	}

	for i, v := range input {
		array.values[i] = v
	}

	for _, i := range missingValueIndices {
		delete(array.values, i)
	}

	return array
}

func ConstructTree[T any](nums *ArrayWithNullableElement[T]) *TreeNode[T] {
	if nums.size == 0 {
		return nil
	}

	if nums.size == 1 {
		val, _ := nums.Get(0)
		return NewTreeNode(val)
	}

	var tree []*TreeNode[T]

	for i := 0; i < nums.size; i++ {
		var treeNode *TreeNode[T]
		num, prs := nums.Get(i)
		if prs {
			treeNode = NewTreeNode(num)
		} else {
			treeNode = nil
		}

		tree = append(tree, treeNode)
	}

	index := 0
	search := 1

	for search < len(tree) {
		node := tree[index]
		if node != nil {
			node.Left = tree[search]
			search = search + 1

			if search < len(tree) {
				node.Right = tree[search]
				search = search + 1
			}
		}

		index = index + 1
	}

	return tree[0]
}
