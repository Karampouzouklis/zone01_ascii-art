# Task 6: Handle Newlines

## Goal
Process `\n` characters to create multiple lines of ASCII art.

## What You'll Learn
- String replacement
- Splitting strings by newlines
- Processing multiple lines

## Steps
1. Split input by newlines:
   ```go
   func renderText(text string, charMap map[rune][]string) {
       // Handle literal \n in input
       text = strings.ReplaceAll(text, "\\n", "\n")
       
       // Split by actual newlines
       lines := strings.Split(text, "\n")
       
       for _, line := range lines {
           if line == "" {
               fmt.Println() // Empty line
           } else {
               renderWord(line, charMap)
           }
       }
   }
   ```

2. Test with newlines:
   ```go
   renderText("Hello\nWorld", charMap)
   renderText("Line1\n\nLine3", charMap) // Empty line in middle
   ```

3. Handle edge cases:
   ```go
   // Empty string
   renderText("", charMap)
   
   // Just newline
   renderText("\n", charMap)
   ```

## Success Criteria
- `\n` creates new lines of ASCII art
- Empty lines are handled correctly
- Multiple consecutive newlines work
- Matches expected output format

## Next: Task 7 - Add basic tests