package asciiartwebstylize

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func PrintAscii1(input string, font string) string {
	//Read the font file------------------------------------
	data, err := os.ReadFile("banners/" + font + ".txt")
	if err != nil {
		log.Println(err)
	}
	//to fix the extra char added in thinkertoy file
	if font == "thinkertoy" {
		data = bytes.ReplaceAll(data, []byte("\r"), []byte(""))
	}
	lines := strings.Split(string(data), "\n")
	if input == "" {
		fmt.Println()
		return ""
	}
	// Matching the input with the font
	P := make([]string, 8)
	for i := 0; i <= 7; i++ {
		for _, char := range input {
			P[i] = P[i] + lines[(((int(char)-32)*9)+1+i)]
		}
	}
	asciiArt := ""
	for j := range P {
		asciiArt += P[j] + "\n"
	}
	return asciiArt
}
func Matching1(i string, font string) string {
	if i == "" {
		return ""
	} else if i == "\n" {
		return "\n"
	} else {
		lines := strings.Split(i, "\n")
		asciiArt := ""
		for _, line := range lines {
			asciiArt += PrintAscii1(line, font)
		}
		// fmt.Print(asciiArt)
		return asciiArt
	}
	return ""
}
