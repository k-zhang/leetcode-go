package easy

func fibonacciNumber(n int) int {
	var f1 = 0
	var f2 = 1

	for i := 0; i < n; i++ {
		var temp = f1
		f1 = f2
		f2 = temp + f1
	}

	return f1
}
