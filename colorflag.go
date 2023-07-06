package asciiART

import (
	"os"
	"strings"
)

var ascii_art struct {
	text         string
	size         int
	text_nocolor string
}

func ColorFlag(args []string, colorname string) {
	if strings.Contains(os.Args[1], "--color=") {
		if len(args) == 3 {
			ascii_art.text = PrintART("", os.Args[2], "standard", colorname)
			Ascii_Print(ascii_art.text)
		} else if len(os.Args) == 4 {
			if os.Args[3] == "shadow" || os.Args[3] == "standard" || os.Args[3] == "thinkertoy" || os.Args[3] == "weird" || os.Args[3] == "refined" {
				ascii_art.text = PrintART("", os.Args[2], os.Args[3], colorname)
				Ascii_Print(ascii_art.text)
			} else {
				ascii_art.text = PrintART(os.Args[2], os.Args[3], "standard", colorname)
				Ascii_Print(ascii_art.text)
			}
		} else if len(os.Args) == 5 {
			ascii_art.text = PrintART(os.Args[2], os.Args[3], os.Args[4], colorname)
			Ascii_Print(ascii_art.text)
		}
	} else {
		PrintError()
	}
}
