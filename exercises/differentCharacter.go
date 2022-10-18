package exercises

import (
	"math"
)

// Strings are always lower case and always differ in just 1 character
func DifferentCharacter(input1, input2 string) string {

	bingoBoard := make(map[string]int)

	for _, character := range input1 {
		bingoBoard[string(character)] += 1
	}

	for _, character := range input2 {
		bingoBoard[string(character)] -= 1
	}

	for key, value := range bingoBoard {
		if math.Abs(float64(value)) == 1 {
			return string(key)
		}
	}
	return "FAIL"
}
