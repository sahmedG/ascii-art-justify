package asciiART

import (
	"fmt"
	"os"
)

func Ascii_Print(Text string) {
	fmt.Print(Text)
}
func ExportFile(outputfile, output string) {
	f, _ := os.Create(outputfile)
	fmt.Fprint(f, output)
}
