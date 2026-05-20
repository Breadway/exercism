package reversestring

func Reverse(input string) string {
	output := ""
	for _, char := range input {
		output = string(char) + output
	}
	return output
}
