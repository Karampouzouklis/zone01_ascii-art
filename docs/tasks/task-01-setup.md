# Task 1: Basic Setup

## Goal
Get a basic Go program running that accepts command line arguments.

## What You'll Learn
- **Go modules**: How Go manages dependencies and project identity
- **Command line arguments**: Reading user input from the terminal
- **Basic Go program structure**: The foundation every Go program needs

## Why This Matters
Every Go project starts with a module - it's like giving your project a name and telling Go how to manage it. Command line arguments let users interact with your program, which is essential for tools like our ASCII art generator.

## Steps
1. **Create `go.mod` file** (this tells Go about your project):
   ```bash
   go mod init ascii-art
   ```
   
   **What this does:** Creates a module named "ascii-art" that Go uses to track your project and any dependencies you might add later.

2. **Create `cmd/ascii-art/main.go`** (your program's entry point):
   ```go
   package main  // This tells Go this is an executable program

   import (
       "fmt"  // For printing output
       "os"   // For accessing command line arguments
   )

   func main() {  // Every Go program starts here
       // os.Args[0] is the program name, os.Args[1] is the first argument
       if len(os.Args) < 2 {
           return  // Exit if no text provided
       }
       
       text := os.Args[1]  // Get the text user wants to convert
       fmt.Println("Input:", text)  // Print it back for now
   }
   ```
   
   **Key concepts:**
   - `package main`: Makes this an executable program (not a library)
   - `os.Args`: A slice containing command line arguments
   - `len()`: Gets the length of a slice
   - `fmt.Println()`: Prints text to the terminal

3. **Test it works:**
   ```bash
   go run ./cmd/ascii-art "Hello"
   ```
   
   **What you should see:** `Input: Hello`
   
   **What's happening:** Go compiles your code and runs it, passing "Hello" as the first argument.

## Success Criteria
- Program runs without errors
- Prints the input text
- Handles empty input correctly

## Next: Task 2 - Read banner files