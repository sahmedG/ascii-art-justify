package asciiART

func MapFont(fontname string) string {
	var fonts = map[string]string{
		"standard":   "../standard.txt",
		"shadow":     "../shadow.txt",
		"thinkertoy": "../thinkertoy.txt",
		"refined": "../refined.txt",
		"weird": "../weird.txt",
		"nicetext": "../nicetext.txt",
	}
	return fonts[fontname]
}
