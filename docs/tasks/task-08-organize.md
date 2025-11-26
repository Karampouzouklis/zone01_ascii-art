# Task 8: Clean Up and Organize

## Goal
Organize your code into clean, readable functions and handle edge cases.

## What You'll Learn
- Code organization
- Error handling best practices
- Function design

## Steps
1. Organize your main.go into clear functions:
   ```go
   package main

   import (
       "fmt"
       "os"
       "strings"
   )

   func main() {
       if len(os.Args) < 2 {
           return
       }
       
       text := os.Args[1]
       
       charMap, err := loadBanner("standard.txt")
       if err != nil {
           fmt.Printf("Error loading banner: %v\n", err)
           return
       }
       
       renderText(text, charMap)
   }

   func loadBanner(filename string) (map[rune][]string, error) {
       // Your banner loading code here
   }

   func renderText(text string, charMap map[rune][]string) {
       // Your text rendering code here
   }
   ```

2. Add proper error handling:
   - Check if banner file exists
   - Handle invalid characters gracefully
   - Validate input parameters

3. Test edge cases:
   ```bash
   go run . ""           # Empty string
   go run . "\n"         # Just newline
   go run . "Hello\n"    # Text with newline
   go run . "Hello\n\nWorld"  # Multiple newlines
   ```

## Success Criteria
- Code is well-organized and readable
- All edge cases handled correctly
- Error messages are helpful
- Output matches the requirements exactly

## Congratulations!
You've built a complete ASCII art generator and learned fundamental Go concepts!