package easy

func isIsomorphic(s, t string) bool {
	charMap := make(map[rune]rune)
	mapped := make(map[rune]bool)

	sArray := []rune(s)
	tArray := []rune(t)

	for index, sRune := range sArray {
		replacement, ok := charMap[sRune]
		if ok {
			if replacement != tArray[index] {
				return false
			}
		} else {
			if _, ok := mapped[tArray[index]]; ok {
				return false
			}

			charMap[sArray[index]] = tArray[index]
			mapped[tArray[index]] = true
		}
	}

	return true
}
