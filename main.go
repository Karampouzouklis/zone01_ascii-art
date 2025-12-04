package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("USAGE: go run . <TEXT> [BANNER]")
		fmt.Println("       BANNER: standard (default), shadow, thinkertoy")
		return
	}

	text := os.Args[1]

	banner := "standard" // Set Default

	if len(os.Args) == 3 {
		banner = strings.ToLower(os.Args[2]) // Here we save the name of a banner style. User should print "standard" or "shadow" or "thinkertoy".
	}

	content, err := readBannerFile(banner, "assets/")
	if err != nil {
		fmt.Println("ERROR: Can not read the banner files!")
		return
	}

	charMap := buildCharMap(content)

	if invalid := validateInput(text, charMap); len(invalid) > 0 {
		// Convert runes to string for display
		invalidChars := make([]string, len(invalid))
		for i, char := range invalid {
			invalidChars[i] = fmt.Sprintf("'%c'", char)
		}
		fmt.Printf("WARNING: Unsupported ASCII Characters %s found and replaced with question mark pattern.\n", strings.Join(invalidChars, ", "))
	}

	renderText(text, charMap)
}
