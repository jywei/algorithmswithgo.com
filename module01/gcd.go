package module01

// GCD is the greatest common divisor
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
		// tmp = a
		// a = b
		// b = tmp%a
	}
	return a
}
