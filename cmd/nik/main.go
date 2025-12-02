package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	//text := os.Args[1]
	//fmt.Println("Input:", text)

	content, err := readBannerFile("assets/standard.txt")
	if err != nil { // Always check for errors!
		fmt.Println("Error:", err)
		return // Exit if we can't read the file
	}
	//fmt.Println("File loaded successfully")

	content = strings.TrimLeft(content, "\n") // Get rid of extra new lines in the beginning of banner file

	charMap := parseCharacters(content)
	//renderChar(' ', charMap)
	renderText("Dj Nik says\nHello World!", charMap)
	//renderText("\n", charMap)
}

func parseCharacters(content string) map[rune][]string {
	characters := strings.Split(content, "\n\n")

	charMap := make(map[rune][]string)

	for i, char := range characters {
		if i < 95 { // ASCII printable characters: 32-126 (95 total)
			ascii := rune(32 + i)              // Convert index to ASCII character
			lines := strings.Split(char, "\n") // Split into 8 lines
			charMap[ascii] = lines             // Store in our map
		}
	}

	return charMap
}

func readBannerFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func renderChar(char rune, charMap map[rune][]string) {
	lines, exist := charMap[char]
	if !exist {
		lines = charMap[' ']
		fmt.Print("Warning: Invalid Character, using space\n")
	}

	for i, line := range lines {
		fmt.Println(strconv.Itoa(i) + ": " + line + "$")
	}

}

func renderWord(word string, charMap map[rune][]string) {
	lines := make([]string, 8)

	for _, char := range word {
		charLines, exist := charMap[char]
		if !exist {
			charLines = charMap[' ']
		}
		for i := 0; i < 8; i++ {
			if i < len(charLines) {
				lines[i] += charLines[i]
			}
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}

}

func renderText(text string, charMap map[rune][]string) {
	text = strings.ReplaceAll(text, "\\n", "\n")

	lines := strings.Split(text, "\n")

	for _, line := range lines {
		if line == "" {
			fmt.Println()
		} else {
			renderWord(line, charMap)
		}
	}
}
