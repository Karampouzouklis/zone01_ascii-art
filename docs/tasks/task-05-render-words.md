# Task 5: Render Words

## Goal
Combine multiple characters horizontally to form words.

## What You'll Learn
- **String concatenation**: Joining strings together
- **Slice initialization**: Creating slices with `make()`
- **2D thinking**: Building output line by line instead of character by character
- **Algorithm design**: Breaking complex problems into simple steps

## Why This Matters
This is where the magic happens! You'll learn to think in two dimensions - combining characters horizontally while managing multiple lines vertically. This kind of spatial reasoning is crucial for many programming problems.

## Steps
1. **Create function to render a word:**
   ```go
   func renderWord(word string, charMap map[rune][]string) {
       // Create 8 empty lines for output (each character is 8 lines tall)
       lines := make([]string, 8)
       
       // For each character in the word
       for _, char := range word {
           pattern := charMap[char]
           if pattern == nil {
               pattern = charMap[' '] // fallback to space
           }
           
           // Add each line of the character to corresponding output line
           for i := 0; i < 8; i++ {
               if i < len(pattern) {
                   lines[i] += pattern[i]  // Concatenate horizontally
               }
           }
       }
       
       // Print all lines
       for _, line := range lines {
           fmt.Println(line)
       }
   }
   ```
   
   **The algorithm explained:**
   1. **Prepare 8 empty strings** - one for each line of the final output
   2. **For each character** in the word:
      - Get its 8-line pattern
      - Add line 0 of the character to line 0 of our output
      - Add line 1 of the character to line 1 of our output
      - ... and so on for all 8 lines
   3. **Print the result** - 8 lines that show all characters side by side
   
   **Key insight:** We build horizontally by concatenating strings, not by printing characters one by one!

2. **Test with simple words:**
   ```go
   renderWord("Hi", charMap)
   fmt.Println() // Empty line between tests
   renderWord("Hello", charMap)
   fmt.Println()
   renderWord("Go!", charMap)  // Test with punctuation
   ```
   
   **What to observe:**
   - Characters should appear side by side, not stacked vertically
   - Each word should be exactly 8 lines tall
   - Spacing should look natural
   
   **Debug tip:** If something looks wrong, print `len(lines)` and `len(pattern)` to check your data!

## Success Criteria
- Characters appear side by side correctly
- All 8 lines are printed for each word
- Spacing between characters looks right

## Next: Task 6 - Handle newlines