package asciiartwebexportfile

import (
	"strings"
)

func Validate(input string) bool {
	if input == "" {
		return false
	}

	inputS := strings.Split(input, "\n")
	for _, line := range inputS {
		line = strings.TrimRight(line, "\r") // remove windows-style carriage return (\r)
		for _, char := range line {
			if char < ' ' || char > '~' {
				return false
			}
		}
	}
	return true
}
func Validatefont(font string) bool {
	if (font == "shadow") || (font == "standard") || (font == "thinkertoy") {
		return true
	}
	return false
}
