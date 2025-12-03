package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [TEXT] [BANNER NAME]")
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

	renderText(text, charMap)
}
