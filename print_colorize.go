package asciiART

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Print_Colorize(color string, input_str string) string{
	if r, g, b := detect_rgb(color); r != -1 && g != -1 && b != -1 {
		r, g, b := detect_rgb(color)
		//fmt.Print(colorizing_string(input_str, RGBColor{r, g, b}))
		return colorizing_string(input_str, RGBColor{r, g, b})
	}
	var Colors = map[string]string{
		"black":   "\033[1;30m%s\033[0m",
		"red":     "\033[1;31m%s\033[0m",
		"green":   "\033[1;32m%s\033[0m",
		"yellow":  "\033[1;33m%s\033[0m",
		"blue":    "\033[1;34m%s\033[0m",
		"magenta": "\033[1;35m%s\033[0m",
		"cyan":    "\033[1;36m%s\033[0m",
		"white":   "\033[1;37m%s\033[0m",
		"orange":  "\033[38;5;208m%s\033[0m",
	}
	if Colors[color] == "" {
		return input_str
	}
	rtnstr := fmt.Sprintf(Colors[color], input_str)
	return rtnstr
}

func detect_rgb(color string) (int, int, int) {
	color_len := len(color)
	if !strings.Contains(color, "rgb(") {
		return -1, -1, -1
	}
	Red_str := ""
	Green_str := ""
	Blue_str := ""
	color_order := 1

	if strings.Count(color, ",") > 3 {
		return -1, -1, -1
	}
	for i := 0; i < color_len; i++ {
		if unicode.IsDigit(rune(color[i])) && color_order == 1 {
			Red_str += string(color[i])
		} else if unicode.IsDigit(rune(color[i])) && color_order == 2 {
			Green_str += string(color[i])
		} else if unicode.IsDigit(rune(color[i])) && color_order == 3 {
			Blue_str += string(color[i])
		}
		if color[i] == ',' {
			color_order++
		}
	}
	Red_val, _ := strconv.Atoi(Red_str)
	Green_val, _ := strconv.Atoi(Green_str)
	Blue_val, _ := strconv.Atoi(Blue_str)

	return Red_val, Green_val, Blue_val
}

// Struct to represent RGB color components
type RGBColor struct {
	R, G, B int
}

// Function to apply color formatting to a string
func colorizing_string(text string, color RGBColor) string {
	// Construct the ANSI escape sequence dynamically
	const EscapeReset = "\033[0m"
	escapeSequence := fmt.Sprintf("\033[38;2;%d;%d;%dm", color.R, color.G, color.B)

	// Apply the color escape sequence to the text
	return escapeSequence + text + EscapeReset
}
