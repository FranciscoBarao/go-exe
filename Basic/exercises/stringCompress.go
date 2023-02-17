package exercises

import "strconv"

/*

Only compress the string if it saves space.

*/
func StringCompress(str string) string {

	compressString, savesSpace, count := "", false, 1

	runeList := []rune(str)

	previous_character := runeList[0]

	count = 0
	for _, character := range runeList {
		if character == previous_character {
			count += 1
		} else {
			if count > 2 {
				savesSpace = true
			}
			if count > 1 {
				compressString += string(previous_character) + strconv.Itoa(count)
			} else {
				compressString += string(previous_character)
			}

			previous_character = character
			count = 1
		}
	}
	if count > 1 {
		compressString += string(previous_character) + strconv.Itoa(count)
	}

	if savesSpace {
		return compressString
	}
	return str
}
