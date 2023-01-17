package fibonacci

func Fibonacci() func(int) []int {
	x := []int{}
	a, b := 0, 1
	return func(n int) []int {
		for i := 0; i < n; i++ {
			f := a
			a, b = b, a+b
			x = append(x, f)
		}
		return x

	}

}
