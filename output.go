package asciiART

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Output(outputfile string, output string, filename string) {

	// 4. Creating a file
	f, _ := os.Create(outputfile)
	res := ""
	if output == "\\n" {
		fmt.Println()
		return
	} else if output == "" {
		return
	}
	input_strs := strings.Split(output, "\\n")
	for _, word := range input_strs {
		if word == "" {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				//* Iterate through each character in the input string
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
							} else {
								res += GetLine(MapFont(filename), MapART(rune(char))+i)
							}
						}
					} else {
						res += GetLine(MapFont(filename), MapART(rune(char))+i)
					}
				}
				if i <= 7 {
					fmt.Fprintln(f, res)
				} else {
					fmt.Fprint(f, res)
					//fmt.Print()
				}
				res = ""
			}
		}
	}

	defer f.Close()
}

func GetLine(filename string, num int) string {

	// * this program scans the banner to get the art
	file, err := os.Open(filename)
	// If there is an error, then handle it
	if err != nil {
		fmt.Println("Error opening file: ", err, ", please enter a valid banner file name")
		os.Exit(1)
	}

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read line by line
	line := ""
	lineCount := 0 // A counter used to stop at specific line
	for scanner.Scan() {
		lineCount++
		// save the line and stop the loop
		if lineCount == num {
			line = scanner.Text()
			break
		}
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		os.Exit(1)
	}
	defer file.Close()
	return line
}
