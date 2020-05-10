package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
// func Reverse(word string) string {
// 	var res string
// 	for i := len(word) - 1; i >= 0; i-- {
// 		res += string(word[i])
// 	}
// 	return res
// }

// Reverse will return the provided word in reverse
// func Reverse(word string) string {
// 	var sb strings.Builder
// 	for i := len(word) - 1; i >= 0; i-- {
// 		sb.WriteByte(word[i])
// 	}
// 	return sb.String()
// }

// Reverse will return the provided word in reverse
// func Reverse(word string) string {
// 	var res string
// 	for i := 0; i < len(word); i++ {
// 		res = string(word[i]) + res
// 	}
// 	return res
// }

// Reverse will return the provided word in reverse with rune
func Reverse(word string) string {
	var res string
	for _, r := range word {
		res = string(r) + res
	}
	return res
}
