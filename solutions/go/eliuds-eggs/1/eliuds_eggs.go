package eliudseggs

func EggCount(displayValue int) int {
	output := 0
	for displayValue > 0 {
		remainder := displayValue % 2
		displayValue /= 2
		output += remainder
	}
	return output
}
