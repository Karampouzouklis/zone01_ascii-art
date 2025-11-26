# Task 2: Read Banner Files

## Goal
Learn to read files in Go and understand the banner format.

## What You'll Learn
- **File reading with `os.ReadFile()`**: How to load data from files
- **Error handling**: Go's way of dealing with things that might go wrong
- **Understanding the banner format**: How ASCII art characters are stored

## Why This Matters
Real programs need to work with external data. Files are the most common way to store and share data. Learning to read files safely (with error handling) is a fundamental skill.

## Steps
1. **Examine the banner files** in the `assets/` directory:
   - `assets/standard.txt`
   - `assets/shadow.txt` 
   - `assets/thinkertoy.txt`
   
   **Take a look inside one:** Each character is represented by exactly 8 lines of ASCII art, and characters are separated by empty lines. This is the "template" we'll use to convert regular text into art.

2. **Create a function to read banner file:**
   ```go
   func readBannerFile(filename string) (string, error) {
       data, err := os.ReadFile(filename)  // Read entire file into memory
       if err != nil {  // Check if something went wrong
           return "", err  // Return empty string and the error
       }
       return string(data), nil  // Convert bytes to string, no error
   }
   ```
   
   **Key concepts:**
   - **Multiple return values**: Go functions can return more than one thing
   - **Error handling**: The `error` type represents things that can go wrong
   - **`nil`**: Go's way of saying "nothing" or "no error"
   - **Type conversion**: `string(data)` converts bytes to a string

3. **Test reading a file:**
   ```go
   content, err := readBannerFile("assets/standard.txt")
   if err != nil {  // Always check for errors!
       fmt.Println("Error:", err)
       return  // Exit if we can't read the file
   }
   fmt.Println("File loaded successfully")
   fmt.Println("First 100 characters:", content[:100])  // Peek at the content
   ```
   
   **What you're learning:**
   - **Error-first programming**: Always handle errors immediately
   - **String slicing**: `content[:100]` gets the first 100 characters
   - **Defensive programming**: Check that things work before proceeding

## Success Criteria
- Can read banner files without errors
- Understand that each character takes 8 lines
- Handle file not found errors

## Next: Task 3 - Parse characters