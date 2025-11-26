# Task 4: Render Single Characters

## Goal
Display ASCII art for individual characters.

## What You'll Learn
- **Map lookups**: Finding values by their keys
- **Conditional logic**: Handling cases where things might not exist
- **Function design**: Creating reusable code blocks
- **Defensive programming**: Making your code handle unexpected situations

## Why This Matters
Before we can render words, we need to master rendering single characters. This teaches you how to safely access data structures and handle edge cases - skills you'll use constantly in programming.

## Steps
1. **Create a function to render one character:**
   ```go
   func renderChar(char rune, charMap map[rune][]string) {
       if pattern, exists := charMap[char]; exists {
           for _, line := range pattern {
               fmt.Println(line)
           }
       }
   }
   ```
   
   **Function anatomy:**
   - **`func`**: Declares a function
   - **`renderChar`**: Function name (use descriptive names!)
   - **`(char rune, charMap map[rune][]string)`**: Parameters the function needs
   - **No return type**: This function just prints, doesn't return anything
   
   **The logic:** Look up the character in our map, and if found, print each line of its ASCII art.

2. Test with different characters:
   ```go
   renderChar('H', charMap)
   renderChar('!', charMap)
   renderChar(' ', charMap) // space character
   ```

3. **Handle characters not in the map** (defensive programming):
   ```go
   func renderChar(char rune, charMap map[rune][]string) {
       pattern, exists := charMap[char]
       if !exists {
           // Fallback to space character for unknown characters
           pattern = charMap[' ']
           fmt.Printf("Warning: Character '%c' not found, using space\n", char)
       }
       for _, line := range pattern {
           fmt.Println(line)
       }
   }
   ```
   
   **Why this matters:**
   - **`!exists`**: The `!` means "not" - so "if not exists"
   - **Fallback strategy**: When something goes wrong, have a reasonable default
   - **User feedback**: Let the user know when something unexpected happens
   - **`%c`**: Printf format for printing a character
   
   **Real-world lesson:** Always think about what happens when your assumptions are wrong!

## Success Criteria
- Can render any single character correctly
- Handles missing characters gracefully
- Output matches expected 8-line format

## Next: Task 5 - Render words