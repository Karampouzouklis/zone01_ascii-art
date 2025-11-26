# Task 3: Parse Characters

## Goal
Convert banner file content into a map of characters to their ASCII art.

## What You'll Learn
- **String splitting with `strings.Split()`**: Breaking text into pieces
- **Maps in Go**: Key-value storage (like a dictionary)
- **ASCII character codes**: How computers represent letters and symbols
- **Runes**: Go's way of handling characters properly

## Why This Matters
Maps are one of the most useful data structures - they let you quickly look up information by a key. Here, we'll map each character (like 'A') to its ASCII art pattern. This makes rendering fast and simple.

## Steps
1. **Split file content by double newlines** (each character block):
   ```go
   characters := strings.Split(content, "\n\n")
   ```
   
   **What this does:** The banner file has each character separated by an empty line (\n\n). This splits the entire file into individual character blocks.
   
   **Try this:** Add `fmt.Println("Found", len(characters), "characters")` to see how many you got!

2. **Create a map to store character patterns:**
   ```go
   charMap := make(map[rune][]string)
   ```
   
   **Understanding this:**
   - **`map[rune][]string`**: A map where keys are characters (runes) and values are slices of strings
   - **`rune`**: Go's type for a single character (handles international characters properly)
   - **`[]string`**: A slice of strings (the 8 lines of ASCII art for each character)
   - **`make()`**: Creates an empty map ready to use

3. **Map each character** (starting from space = ASCII 32):
   ```go
   for i, char := range characters {
       if i < 95 { // ASCII printable characters: 32-126 (95 total)
           ascii := rune(32 + i)  // Convert index to ASCII character
           lines := strings.Split(char, "\n")  // Split into 8 lines
           charMap[ascii] = lines  // Store in our map
       }
   }
   ```
   
   **The magic explained:**
   - **ASCII 32**: The space character (first printable character)
   - **ASCII 126**: The ~ character (last printable character)
   - **`32 + i`**: Maps array index 0 to space (32), index 1 to ! (33), etc.
   - **`rune(32 + i)`**: Converts the number to an actual character
   
   **Why 95 characters?** From space (32) to ~ (126) = 126 - 32 + 1 = 95 characters

4. **Test by printing a character:**
   ```go
   if pattern, exists := charMap['A']; exists {
       fmt.Println("Character 'A':")
       for _, line := range pattern {
           fmt.Println(line)
       }
   } else {
       fmt.Println("Character 'A' not found!")
   }
   ```
   
   **What's happening:**
   - **`charMap['A']`**: Look up the letter 'A' in our map
   - **`pattern, exists`**: Get both the value AND whether it was found
   - **`range pattern`**: Loop through each of the 8 lines
   - **`_`**: We don't need the index, so we ignore it with underscore
   
   **Try different characters:** Test with 'B', '!', ' ' (space), or '5'

## Success Criteria
- Successfully parse all characters from banner file
- Can look up any character and get its 8-line pattern
- Handle characters that don't exist

## Next: Task 4 - Render single characters