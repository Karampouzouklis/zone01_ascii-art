# ASCII Art Generator

![Go Version](https://img.shields.io/badge/Go-1.24+-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Tests](https://github.com/karampouzouklis/zone01_ascii-art/actions/workflows/test.yml/badge.svg)


A command-line tool that converts text into ASCII art using different banner styles!

## Features

- Convert any text to ASCII art
- Multiple banner styles: `standard`, `shadow`, `thinkertoy`
- Remote banner file loading with local fallback
- Handle newlines and special characters
- Custom question mark pattern for unsupported characters
- Automatic fallback to standard banner if style not found
- Comprehensive error handling

## Example Output

```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
```

## Installation

1. Make sure you have Go installed (1.24+)
2. Clone this repository
3. Navigate to the project directory

## Usage

```bash
# Basic usage with default (standard) banner
go run . "Hello World"

# Specify a banner style
go run . "Hello World" standard
go run . "Hello World" shadow
go run . "Hello World" thinkertoy

# Handle newlines
go run . "Hello\nWorld"

# Empty lines
go run . "Line1\n\nLine3"
```

## Project Structure

```
ascii-art/
├── main.go          # Application entry point
├── ascii.go         # Core ASCII art functions
├── ascii_test.go    # Comprehensive test suite
├── go.mod           # Go module definition
├── assets/          # Banner template files
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── docs/            # Learning materials and requirements
└── README.md        # This file
```

## Functions

- `readBannerFile()` - Reads banner template files with fallback
- `buildCharMap()` - Parses banner content into character mappings
- `renderText()` - Handles text processing and newlines
- `renderWord()` - Renders individual words as ASCII art

## Testing

Run the comprehensive test suite:

```bash
go test -v
```

Tests cover:
- File reading and error handling
- Character map building
- Text rendering with various inputs
- Edge cases and error conditions

## Error Handling

- Banner files are first attempted from remote URL, then local fallback
- Invalid banner names automatically fallback to `standard`
- Unsupported characters display a custom question mark pattern
- Graceful handling of empty inputs
- Clear error messages for file reading issues

## Learning Journey

This project demonstrates:
- **Go Basics**: modules, functions, variables, error handling
- **File Operations**: reading files, handling missing files
- **Data Structures**: maps, slices, string manipulation
- **Testing**: unit tests, table-driven tests, output capture
- **Code Organization**: clean, maintainable code structure

## Note

This is a personal learning project created to demonstrate Go programming skills. It's not actively maintained or open for contributions, but feel free to fork it for your own learning!