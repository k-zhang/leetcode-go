package easy

import (
	"github.com/golang-collections/collections/stack"
)

func reverseSlice(s []rune) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

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

	var runes []rune
	for workStack.Len() > 0 {
		r := workStack.Pop().(rune)
		runes = append(runes, r)
	}

	reverseSlice(runes)

	return string(runes)
}
