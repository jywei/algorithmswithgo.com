package module01

// Sum will sum up all of the numbers passed
// in and return the result
// func Sum(numbers []int) int {
// 	sum := 0
// 	for _, val := range numbers {
// 		sum += val
// 	}
// 	return sum
// }

// Sum will sum up all of the numbers passed
// with recursion approach
func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum(numbers[1:])
}
