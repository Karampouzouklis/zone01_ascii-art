package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	////////////////////////////////
	//   # Task 1: Basic Setup   //
	//////////////////////////////

	// The below "if" function check input arguments and save the text and banner name.

	if len(os.Args) < 2 { // Check for empty input.
		fmt.Println("ERROR\n", "Example use: go run main.go [TEXT] [BANNER NAME]")
		return
	}
	argStr := os.Args[1]
	var banner string

	if len(os.Args) == 2 { // if input arguments is only [TEXT] then the banner style will be "standart.txt" automaticly.
		banner = "standart"
	} else if len(os.Args) == 3 {
		banner = strings.ToLower(os.Args[2]) // Here we save the name of a banner style. User should print "standard" or "shadow" or "thinkertoy".
	} else {
		fmt.Print("ERROR.\n", "Usage: go run . [STRING] [BANNER] || Example: go run . \"test\" standard or shadow or thinkertoy ")
		return
	}

	sepArgs := strings.Split(argStr, "\\n")

	////////////////////////////////
	//# Task 2: Read Banner Files//
	//////////////////////////////

	file, err := os.ReadFile("assets/" + banner + ".txt") // read the "banner name. Marge it with ".txt" and read the file
	if err != nil {
		fmt.Println("ERROR\n", banner+"Does not exist.")
		return
	}

	lines := strings.Split(string(file), "\n")

	PrintAsciiArt(sepArgs, lines)
}

func PrintAsciiArt(word []string, textfile []string) {
}
