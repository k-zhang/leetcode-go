package easy

import (
	"errors"
	"math"
	"strconv"
)

func reverseSliceWithStart(s []rune, start int) {
	for i, j := start, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseInteger(x int) (int, error) {
	s := strconv.Itoa(x)
	chars := []rune(s)
	if len(chars) == 1 {
		return x, nil
	}

	var start int

	if chars[0] == '-' {
		start = 1
	} else {
		start = 0
	}

	reverseSliceWithStart(chars, start)

	resultString := string(chars)

	r, err := strconv.Atoi(resultString)
	if err != nil {
		return -1, errors.New("failed")
	}

	return r, nil
}

func reverseInteger2(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x = x / 10

		if rev > math.MaxInt/10 || (rev == math.MaxInt/10 && pop > 7) {
			return 0
		}

		if rev < math.MinInt/10 || (rev == math.MinInt/10 && pop < -1) {
			return 0
		}

		rev = rev*10 + pop
	}

	return rev
}
