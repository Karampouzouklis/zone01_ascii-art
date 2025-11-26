# ASCII Art Generator

Learn Go by building a command-line tool that converts text into ASCII art!

## What You'll Build

A program that takes text input and creates ASCII art like this:

```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
```

## Quick Start

1. Make sure you have Go installed
2. Start with Task 1 in `docs/tasks/`
3. Work through tasks 1-8 in order
4. Each task builds on the previous one

## What You'll Learn

- **Go Basics**: modules, functions, variables
- **File Operations**: reading files, handling errors
- **Data Structures**: maps, slices, strings
- **String Processing**: splitting, joining, manipulation
- **Testing**: writing and running tests
- **Code Organization**: clean, readable code

## Learning Path

| Task | Topic | What You'll Do |
|------|-------|----------------|
| 1 | Setup | Create basic Go program with command line args |
| 2 | File Reading | Load banner template files |
| 3 | Parsing | Convert file content into character maps |
| 4 | Single Chars | Render individual ASCII art characters |
| 5 | Words | Combine characters to make words |
| 6 | Newlines | Handle `\n` for multiple lines |
| 7 | Testing | Write basic tests for your code |
| 8 | Organization | Clean up and structure your code |

## Project Structure

Follows Go best practices while staying simple:
```
ascii-art/
├── cmd/ascii-art/   # Application entry point
│   └── main.go      # Your main program
├── assets/          # Banner template files
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── test/            # Test files
├── go.mod           # Go module definition
└── docs/            # Learning tasks and requirements
```

## Getting Started

Go to `docs/tasks/task-01-setup.md` and start coding!

## Final Usage

When complete, your program will work like this:
```bash
go run ./cmd/ascii-art "Hello"
go run ./cmd/ascii-art "Hello\nWorld"
```