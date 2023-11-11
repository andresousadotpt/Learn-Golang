package mystrings

// we capitalized R because we want to export this function and able to use it in my main package
func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
