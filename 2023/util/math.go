package util

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func LCM(x ...int) int {
	if len(x) == 1 {
		return x[0]
	} else if len(x) > 2 {
		return LCM(x[0], LCM(x[1:]...))
	}

	return x[0] * x[1] / gcd(x[0], x[1])
}
