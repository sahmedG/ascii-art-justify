package asciiART

import (
	"fmt"
	"os"
)

func PrintError() {
	fmt.Println(`Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --color=<color> --output=<outputfile.txt> <letters to be colored> "something" shadow`)
	os.Exit(1)
}
