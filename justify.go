package asciiART

import (
	"fmt"
	"strings"
)

var termWidth int

func init() {
	termWidth = GetTermWidth()
}

func Aliging(letters_to_be_colored, text, font, align, color string) string {
	if align == "left" || align == "right" || align == "center" {
		return RightLeft(letters_to_be_colored, text, MapFont(font), align, color)
	} else if align == "justify" {
		string_beg_end, _ := ContainsString(letters_to_be_colored, text)
		lines := strings.Split(text, "\\n")
		for _, v := range lines {
			Justify(string_beg_end,letters_to_be_colored, text, v, MapFont(font), color)
		}
	}
	return ""
}
func RightLeft(letters_to_be_colored, text, font, align, color string) string {
	res := ""
	newres := ""
	string_beg_end, _ := ContainsString(letters_to_be_colored, text)
	args := strings.Split(text, "\\n")
	for _, word := range args {
		for i := 0; i <= 8; i++ {
			str_len := len(word)
			for idx := 0; idx < str_len; idx++ {
				char := word[idx]
				//* Check for the backslash (since this is taken from os arguments you have to read it like this '\\') and then read the letter after it
				if char == '\\' {
					if idx < len(word)-1 {
						//* Apply tab if 't' appears
						if word[idx+1] == 't' {
							res += "\t"
							idx++
							continue
						}
					}
				}
				if IsInRange(string_beg_end, idx) || letters_to_be_colored == "" {
					// Start printing the colored letler ART
					res += PrintFileLine(MapART(rune(char))+i, (font), color)
				} else {
					// Start printing the letler ART in default color
					res += PrintFileLine(MapART(rune(char))+i, (font), "")
				}
				// if i == 0 {
				// 	size = len(res)
				// }
			}
			if align == "left" {
				newres += res
				newres += "\n"
			} else if align == "right" {
				newres += printSpaces(termWidth-len(res)) + res
				newres += "\n"
			} else if align == "center" {
				newres += printSpaces((termWidth-len(res))/2) + res
				newres += "\n"
			}
			res = ""
		}
	}
	return newres
}

func Justify(string_beg_end[][]int,letters_to_be_colored, text, words, font, color string) {
	sws := SplitWhiteSpacesAWESOME(words)
	ar := make([][]string, len(sws))
	j := 0
	container := ""
	for i := 0; i < 8; i++ {
		j = 0
		for k, letter := range words {
			if letter != ' ' {
				if IsInRange(string_beg_end, k) || letters_to_be_colored == "" {
					// Start printing the colored letler ART
					container += PrintFileLine(MapART((letter))+i, (font), color)
				} else {
					// Start printing the letler ART in default color
					container += PrintFileLine(MapART((letter))+i, (font), "")
				}
				//container += PrintFileLine((MapART(letter) + i), font, "")
			} else if letter == ' ' && container != "" {
				ar[j] = append(ar[j], container)
				container = ""
				j++
			}
		}
		ar[j] = append(ar[j], container)
		container = ""
	}
	textLen := 0
	for _, arOfStr := range ar {
		textLen += len(arOfStr[0])
	}
	if len(sws) == 1 {
		Ascii_Print(RightLeft(letters_to_be_colored, text, font, "left", color))
	} else {
		numSpaces := (termWidth - textLen) / (len(sws) - 1)
		for i := 0; i < 8; i++ {
			for k, v := range ar {
				fmt.Print(v[i])
				if k != len(ar)-1 {
					fmt.Print(printSpaces(numSpaces))
				}
			}
			fmt.Println()
		}
	}
}

func printSpaces(num int) string {
	a := ""
	for i := 1; i <= num; i++ {
		a += " "
	}
	return a
}

func GetTermWidth() int {
	// cmd := exec.Command("stty", "size")
	// cmd.Stdin = os.Stdin
	// out, err := cmd.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// out = out[:len(out)-1]
	// tput, _ := strconv.ParseInt(string(out[3:]), 10, 32)
	// tput2 := int(tput)
	// return tput2
	return 186
}

func SplitWhiteSpacesAWESOME(str string) []string {

	nbw := 0
	var prel rune
	prel = ' '
	word := ""

	for _, v := range str {
		if (v != ' ' && v != '\t' && v != '\n') && (prel == ' ' || prel == '\t' || prel == '\n') {
			nbw++
		}
		prel = v
	}

	ar := make([]string, nbw)
	a := 0
	for _, v := range str {
		if v != ' ' && v != '\t' && v != '\n' {
			word = word + string(v)
		}
		if v == ' ' || v == '\t' || v == '\n' && word != "" {
			a++
			ar[a-1] = word
			word = ""
		}

	}

	if word != "" {
		ar[a] = word
	}
	return ar
}
