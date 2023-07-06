package asciiART

import "fmt"

func Print_Each_Rune_Line(letters_to_be_colored string, str string, fontname string, color string) (string,int){
	res := ""
	size := 0
	string_beg_end, _ := ContainsString(letters_to_be_colored, str)
	// if there is parts to be colored

	//* Iterate through eight lines
	for i := 0; i < 8; i++ {
		//* Iterate through each character in the input string
		str_len := len(str)
		for idx := 0; idx < str_len; idx++ {
			char := str[idx]
			//* Check for the backslash (since this is taken from os arguments you have to read it like this '\\') and then read the letter after it
			if char == '\\' {
				if idx < len(str)-1 {
					//* Apply tab if 't' appears
					if str[idx+1] == 't' {
						fmt.Print("\t")
						idx++
						continue
					}
				}
			}
			if IsInRange(string_beg_end, idx) || letters_to_be_colored == "" {
				// Start printing the colored letler ART
				res += PrintFileLine(MapART(rune(char))+i, MapFont(fontname), color)
			} else {
				// Start printing the letler ART in default color
				res += PrintFileLine(MapART(rune(char))+i, MapFont(fontname), "")
			}
			if i == 0 {
				size = len(res)
			}
		}
		res += "\n"
	}
	return res,size
}
