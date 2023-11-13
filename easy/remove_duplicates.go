package easy

import (
	"github.com/golang-collections/collections/stack"
)

func removeDuplicates(s string) string {
	if len(s) == 1 {
		return s
	}

	workStack := stack.New()

	for _, r := range s {
		if workStack.Len() != 0 && workStack.Peek() == r {
			workStack.Pop()
		} else {
			workStack.Push(r)
		}
	}

	sb := ""
	stackSize := workStack.Len()
	for i := 0; i < stackSize; i++ {
		c := workStack.Pop().(int32)
		sb = string(c) + sb
	}

	return sb
}
