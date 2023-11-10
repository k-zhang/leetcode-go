package easy

import "math"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	rev := 0
	y := x

	for y != 0 {
		pop := y % 10
		y = y / 10

		if rev > math.MaxInt/10 || (rev == math.MaxInt/10 && pop > 7) {
			return false
		}

		rev = rev*10 + pop
	}

	return rev == x
}
