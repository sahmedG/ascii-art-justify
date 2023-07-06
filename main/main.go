package main

import (
	"asciiART"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var ascii_art struct {
		text string
		size int
		text_nocolor string
	}
	// Flags
	colorname := flag.String("color", "", "coloring")
	outputfile := flag.String("output", "", "output")
	justify := flag.String("align","","aliging")
	flag.Parse()
	if *colorname != "" && *outputfile == "" && *justify == "" && strings.Contains(os.Args[1], "--color=") {
		if len(os.Args) == 3 {
			ascii_art.text,_ = asciiART.PrintART("", os.Args[2], "standard", *colorname)
			Ascii_Print(ascii_art.text)
		} else if len(os.Args) == 4 {
			if os.Args[3] == "shadow" || os.Args[3] == "standard" || os.Args[3] == "thinkertoy" || os.Args[3] == "weird" || os.Args[3] == "refined" {
				ascii_art.text, _ = asciiART.PrintART("", os.Args[2], os.Args[3], *colorname)
				Ascii_Print(ascii_art.text)
			} else {
				ascii_art.text ,_ = asciiART.PrintART(os.Args[2], os.Args[3], "standard", *colorname)
				Ascii_Print(ascii_art.text)
			}
		} else if len(os.Args) == 5 {
			ascii_art.text ,_ = asciiART.PrintART(os.Args[2], os.Args[3], os.Args[4], *colorname)
			Ascii_Print(ascii_art.text)
		}
		return

	} else if *colorname != "" && *outputfile != "" && (strings.Contains(os.Args[1], "--color=") && strings.Contains(os.Args[2], "--output=") || strings.Contains(os.Args[2], "--color=") && strings.Contains(os.Args[1], "--output=")) {
		if len(os.Args) == 4 {
			ascii_art.text,_ = asciiART.PrintART("", os.Args[3], "standard", *colorname)
			Ascii_Print(ascii_art.text)
			ascii_art.text_nocolor,_ = asciiART.PrintART("", os.Args[3], "standard", "")
			ExportFile(*outputfile,ascii_art.text_nocolor)
		} else if len(os.Args) == 5 {
			if os.Args[4] == "shadow" || os.Args[4] == "standard" || os.Args[4] == "thinkertoy" || os.Args[4] == "weird" || os.Args[4] == "refined" {
				ascii_art.text,_ = asciiART.PrintART("", os.Args[3], os.Args[4], *colorname)
				Ascii_Print(ascii_art.text)
				ascii_art.text_nocolor,_ = asciiART.PrintART("", os.Args[3], os.Args[4], "")
				ExportFile(*outputfile,ascii_art.text_nocolor)
			} else {
				ascii_art.text,_ = asciiART.PrintART(os.Args[3], os.Args[4], "standard", *colorname)
				Ascii_Print(ascii_art.text)
				ascii_art.text_nocolor,_ = asciiART.PrintART("", os.Args[4], "standard", "")
				ExportFile(*outputfile,ascii_art.text_nocolor)
			}
		} else if len(os.Args) == 6 {
			ascii_art.text,_ = asciiART.PrintART(os.Args[3], os.Args[4], os.Args[5], *colorname)
			Ascii_Print(ascii_art.text)
			ascii_art.text_nocolor,_ = asciiART.PrintART(os.Args[3], os.Args[4], os.Args[5], "")
			ExportFile(*outputfile,ascii_art.text_nocolor)
		}
		return
	} else if *colorname == "" && *justify == "" && *outputfile != "" && strings.Contains(os.Args[1], "--output=") {
		if len(os.Args) == 3 {
			ascii_art.text,_ = asciiART.PrintART("", os.Args[2], "standard", "")
			ExportFile(*outputfile, ascii_art.text)
		} else if len(os.Args) == 4 {
			ascii_art.text,_ = asciiART.PrintART("", os.Args[2], os.Args[3], "")
			ExportFile(*outputfile, ascii_art.text)
		}
		return
	} else if *colorname == "" && *justify != "" && *outputfile == "" && strings.Contains(os.Args[1], "--align="){
		if len(os.Args) == 3 {
			asciiART.Aliging(os.Args[2],"standard",*justify)
		} else if len(os.Args) == 4 {
			asciiART.Aliging(os.Args[2],"standard",*justify)
		}
		return
	} else {
		if len(os.Args) == 2 {
			ascii_art.text ,_ = asciiART.PrintART("", os.Args[1], "standard", "")
			Ascii_Print(ascii_art.text)
			return
		} else if len(os.Args) == 3 {
			ascii_art.text,_ = asciiART.PrintART("", os.Args[1], os.Args[2], "")
			Ascii_Print(ascii_art.text)
			return
		}
	}

	fmt.Println(`Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --color=<color> --output=<outputfile.txt> <letters to be colored> "something" shadow`)
}

func Ascii_Print(Text string) {
	fmt.Print(Text)
}
func ExportFile(outputfile, output string) {
	f, _ := os.Create(outputfile)
	fmt.Fprint(f, output)
}