package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// The below "if" function check input arguments and save the text and banner name.

	if len(os.Args) < 2 { // Check for empty input.
		fmt.Println("ERROR\n", "Example use: go run main.go [TEXT] [BANNER NAME]")
	} else if len(os.Args) == 2 { // if input arguments is only [TEXT] then the banner style will be "standart.txt" automaticly.
		banner := "standart"
		fmt.Println(banner)
	} else if len(os.Args) == 3 {
		banner := os.Args[2] // Here we save the name of a banner style. User should print "standard" or "shadow" or "thinkertoy".
		fmt.Println(banner)
	} else {
		fmt.Print("ERROR\n", "Cannot read the input arguments")
	}

	argStr := strings.Split(os.Args[1], "\n")

	fmt.Println(argStr)
}
