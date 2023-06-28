package asciiART

/*
This function will test if the letter to be printed is to be colored or not
inputs:

	letters:		letters to be colored
	input_letter:	the letter to be tested

output:

	bool			(true or false)
*/
func ContainsLetter(letters string, input_letter string) bool {
	for _, letter := range letters {
		if letter == rune(input_letter[0]) {
			return true
		}
	}

	return false
}
