package asciiART

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var termWidth int

func init() {
	termWidth = GetTermWidth()
}

func Aliging(text, font, align, color string) string {
	if align == "left" || align == "right" || align == "center" {
		return RightLeft(text, MapFont(font), align, color)
	} else if align == "justify" {
		lines := strings.Split(text, "\\n")
		for _, v := range lines {
			Justify(text, v, MapFont(font), color)
			return ""
		}
	}
	return ""
}
func RightLeft(text, font, align, color string) string {
	res := ""
	newres := ""
	args := strings.Split(text, "\\n")
	for _, word := range args {
		for i := 0; i <= 8; i++ {
			for _, letter := range word {
				res += PrintFileLine(MapART(letter)+i, font, color)
			}
			if align == "left" {
				newres += res
				if i < 7 {
					newres += "\n"
				}
			} else if align == "right" {
				newres += printSpaces(termWidth-len(res)) + res
				if i < 7 {
					newres += "\n"
				}
			} else if align == "center" || align == "centre" {
				newres += printSpaces((termWidth-len(res))/2) + res
				if i < 7 {
					newres += "\n"
				}
			}
			res = ""
		}
	}
	return newres
}

func Justify(text, words, font, color string) {
	sws := SplitWhiteSpacesAWESOME(words)
	ar := make([][]string, len(sws))
	j := 0
	container := ""
	for i := 0; i < 8; i++ {
		j = 0
		for _, letter := range words {
			if letter != ' ' {
				container += PrintFileLine((MapART(letter) + i), font, "")
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
		Ascii_Print(RightLeft(text, font, "left", color))
	} else {
		numSpaces := (termWidth - textLen) / (len(sws) - 1)
		for i := 0; i < 8; i++ {
			for k, v := range ar {

				fmt.Print(Print_Colorize(color, v[i]))
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
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	out = out[:len(out)-1]
	tput, _ := strconv.ParseInt(string(out[3:]), 10, 32)
	tput2 := int(tput)
	return tput2
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
