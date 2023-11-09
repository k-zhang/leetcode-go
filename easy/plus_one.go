package easy

func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+1 < 10 {
			digits[i] += 1
			break
		} else {
			digits[i] = 0
		}
	}

	if digits[0] == 0 {
		var newDigits = append([]int{1}, digits...)
		digits = newDigits
	}

	return digits
}
