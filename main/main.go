package main

import (
	"asciiART"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Flags
	colorname := flag.String("color", "", "coloring")
	outputfile := flag.String("output", "", "output")
	justify := flag.String("align","","aliging")
	flag.Parse()
	if *colorname != "" && *outputfile == "" && *justify == "" && strings.Contains(os.Args[1], "--color=") {
		if len(os.Args) == 3 {
			asciiART.PrintART("", os.Args[2], "standard", *colorname)
		} else if len(os.Args) == 4 {
			if os.Args[3] == "shadow" || os.Args[3] == "standard" || os.Args[3] == "thinkertoy" || os.Args[3] == "weird" || os.Args[3] == "refined" {
				asciiART.PrintART("", os.Args[2], os.Args[3], *colorname)
			} else {
				asciiART.PrintART(os.Args[2], os.Args[3], "standard", *colorname)
			}
		} else if len(os.Args) == 5 {
			asciiART.PrintART(os.Args[2], os.Args[3], os.Args[4], *colorname)
		}
		return

	} else if *colorname != "" && *outputfile != "" && (strings.Contains(os.Args[1], "--color=") && strings.Contains(os.Args[2], "--output=") || strings.Contains(os.Args[2], "--color=") && strings.Contains(os.Args[1], "--output=")) {
		if len(os.Args) == 4 {
			asciiART.PrintART("", os.Args[3], "standard", *colorname)
			asciiART.Output(*outputfile, os.Args[3], "standard")
		} else if len(os.Args) == 5 {
			if os.Args[4] == "shadow" || os.Args[4] == "standard" || os.Args[4] == "thinkertoy" || os.Args[4] == "weird" || os.Args[4] == "refined" {
				asciiART.PrintART("", os.Args[3], os.Args[4], *colorname)
				asciiART.Output(*outputfile, os.Args[3], os.Args[4])
			} else {
				asciiART.PrintART(os.Args[3], os.Args[4], "standard", *colorname)
				asciiART.Output(*outputfile, os.Args[4], "standard")
			}
		} else if len(os.Args) == 6 {
			asciiART.PrintART(os.Args[3], os.Args[4], os.Args[5], *colorname)
			asciiART.Output(*outputfile, os.Args[4], os.Args[5])
		}
		return
	} else if *colorname == "" && *justify == "" && *outputfile != "" && strings.Contains(os.Args[1], "--output=") {
		if len(os.Args) == 3 {
			asciiART.Output(*outputfile, os.Args[2], "standard")
		} else if len(os.Args) == 4 {
			asciiART.Output(*outputfile, os.Args[2], os.Args[3])
		}
		return
	} else if *colorname == "" && *justify != "" && *outputfile == "" && strings.Contains(os.Args[1], "--output=") {
		if len(os.Args) == 3 {
			asciiART.Aliging(os.Args[2],"standard",*justify)
		} else if len(os.Args) == 4 {
			asciiART.Aliging(os.Args[2],os.Args[3],*justify)
		}
		return
	} else {
		if len(os.Args) == 2 {
			asciiART.PrintART("", os.Args[1], "standard", "")
			return
		} else if len(os.Args) == 3 {
			asciiART.PrintART("", os.Args[1], os.Args[2], "")
			return
		}
	}

	fmt.Println(`Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --color=<color> --output=<outputfile.txt> <letters to be colored> "something" shadow`)
}
