package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// No input → no output (as in subject examples)
	if len(os.Args) < 2 {
		return
	}

	// Join all arguments so "Hello There" works too
	text := strings.Join(os.Args[1:], " ")

	// Load banner (font) and build character map
	charMap, err := loadBanner("assets/standard.txt")
	if err != nil {
		// Print errors to stderr so stdout remains only ASCII art
		fmt.Fprintln(os.Stderr, "Error loading banner:", err)
		return
	}

	// Render the full text (supports \n, empty lines, etc.)
	renderText(text, charMap)
}

// loadBanner reads the banner file and builds the rune→ASCII-art map
func loadBanner(filename string) (map[rune][]string, error) {
	content, err := readBannerFile(filename)
	if err != nil {
		return nil, err
	}

	charMap := buildCharMap(content)
	if len(charMap) == 0 {
		return nil, fmt.Errorf("character map is empty")
	}

	return charMap, nil
}

// readBannerFile reads the whole banner file into a string
func readBannerFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// buildCharMap converts banner content into a map from rune to its 8-line pattern
func buildCharMap(content string) map[rune][]string {
	// Normalize Windows CRLF to Unix LF
	content = strings.ReplaceAll(content, "\r\n", "\n")

	// Split into character blocks separated by one empty line
	charBlocks := strings.Split(content, "\n\n")

	charMap := make(map[rune][]string)

	// ASCII printable range: 32 (' ') to 126 ('~') → 95 characters
	for i, block := range charBlocks {
		if i >= 95 {
			break
		}

		asciiChar := rune(32 + i)
		lines := strings.Split(block, "\n")

		// Ensure we have at least 8 lines for this character
		if len(lines) < 8 {
			continue
		}
		if len(lines) > 8 {
			lines = lines[:8]
		}

		charMap[asciiChar] = lines
	}

	return charMap
}

// renderWord renders a single line (word/string without '\n') as ASCII art
func renderWord(word string, charMap map[rune][]string) {
	const height = 8

	// Prepare 8 empty output lines
	lines := make([]string, height)

	// For each character in the word
	for _, char := range word {
		pattern, exists := charMap[char]
		if !exists {
			// Fallback to space if character is missing
			pattern, exists = charMap[' ']
			if !exists {
				// If even space is missing, skip this character
				continue
			}
		}

		// Concatenate each row of this character to the corresponding output row
		for i := 0; i < height; i++ {
			if i < len(pattern) {
				lines[i] += pattern[i]
			}
		}
	}

	// Print the final 8-line result
	for _, line := range lines {
		fmt.Println(line)
	}
}

// renderText handles literal "\n", real newlines and empty lines
func renderText(text string, charMap map[rune][]string) {
	// If text is completely empty → no output (matches subject)
	if text == "" {
		return
	}

	// Turn literal "\n" into real newlines (so "Hello\\nWorld" works)
	text = strings.ReplaceAll(text, "\\n", "\n")

	// Split by actual newline characters
	lines := strings.Split(text, "\n")

	// If the text ends with a newline, strings.Split will give a trailing ""
	// We handle that so that "\n" prints one empty line, not two
	for i, line := range lines {
		isLast := i == len(lines)-1
		if isLast && line == "" && strings.HasSuffix(text, "\n") {
			// Skip the trailing empty element created by Split
			break
		}

		if line == "" {
			// Logical empty line → print a blank line
			fmt.Println()
		} else {
			renderWord(line, charMap)
		}
	}
}
