package asciiART

import (
	"fmt"
	"strings"
)

//* this function prints the art while handling newlines

func PrintART(letters_to_be_colored string, str string, fontname string, color string) {
	//* in the case newlines and only newlines were detected, the program prints a new line and exits
	if str == "\\n" {
		fmt.Println()
		return
	}
	input_strs := strings.Split(str, "\\n")
	for _, word := range input_strs {
		if word == "" {
			fmt.Println()
		} else {
			Print_Each_Rune_Line(letters_to_be_colored, word, fontname, color)
		}
	}
}
