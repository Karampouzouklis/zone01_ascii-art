package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// readBannerFile reads the whole banner file from remote URL first, then local fallback
func readBannerFile(filename string, path string) (string, error) {
	// Try remote URL first
	remoteURL := "https://platform.zone01.gr/api/content/root/public/subjects/ascii-art/" + filename + ".txt"
	resp, err := http.Get(remoteURL)
	if err == nil && resp.StatusCode == 200 {
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		if err == nil {
			return string(data), nil
		}
	}

	// Fallback to local file
	data, err := os.ReadFile(path + filename + ".txt")
	if err != nil && filename != "standard" {
		fmt.Printf("WARNING: Banner '%s' not found, using standard\n", filename)
		// Try remote standard first
		resp, err := http.Get("https://platform.zone01.gr/api/content/root/public/subjects/ascii-art/standard.txt")
		if err == nil && resp.StatusCode == 200 {
			defer resp.Body.Close()
			data, err := io.ReadAll(resp.Body)
			if err == nil {
				return string(data), nil
			}
		}
		// Fallback to local standard
		data, err = os.ReadFile(path + "standard.txt")
		if err != nil {
			return "", err
		}
	}

	return string(data), nil
}

func buildCharMap(content string) map[rune][]string {
	// Handle both Windows (\r\n) and Unix (\n) line endings
	content = strings.ReplaceAll(content, "\r\n", "\n")
	// Remove leading newlines that exist in banner files
	content = strings.TrimLeft(content, "\n")
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

// renderWord renders a single line (word/string without '\n') as ASCII art
func renderWord(word string, charMap map[rune][]string) {
	const height = 8

	// Prepare 8 empty output lines
	lines := make([]string, height)

	//invalidChar := []string{" I ", " n ", " v ", " a ", " l ", " i ", " d ", " ! "}
	invalidChar := []string{" __    ", "|***\\  ", "   )*| ", "  /*/  ", " |*|   ", " (*)   ", "       ", "       "}

	// For each character in the word
	for _, char := range word {
		pattern, exists := charMap[char]
		if !exists {
			// Fallback if character is missing
			pattern = invalidChar //charMap[' ']
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

func validateInput(text string, charMap map[rune][]string) []rune {
	var invalid []rune
	for _, char := range text {
		if char != '\n' && char != '\\' { // Skip newlines and escape chars
			if _, exists := charMap[char]; !exists {
				// Check if we already added this character
				found := false
				for _, inv := range invalid {
					if inv == char {
						found = true
						break
					}
				}
				if !found {
					invalid = append(invalid, char)
				}
			}
		}
	}
	return invalid
}
