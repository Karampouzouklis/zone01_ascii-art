package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// Test readBannerFile with valid file
func TestReadBannerFile_ValidFile(t *testing.T) {
	content, err := readBannerFile("standard", "assets/")
	if err != nil {
		t.Fatalf("Failed to read banner file: %v", err)
	}
	if len(content) == 0 {
		t.Fatal("Banner file is empty")
	}
}

// Test readBannerFile with invalid file (should fallback to standard)
func TestReadBannerFile_InvalidFile(t *testing.T) {
	// Capture stdout to check warning message
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	content, err := readBannerFile("nonexistent", "assets/")

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if err != nil {
		t.Fatalf("Should fallback to standard, got error: %v", err)
	}
	if len(content) == 0 {
		t.Fatal("Should have fallback content")
	}
	if !strings.Contains(output, "WARNING") {
		t.Error("Should print warning message")
	}
}

// Test buildCharMap with mock data
func TestBuildCharMap_BasicParsing(t *testing.T) {
	// Mock content: space (32) and exclamation (33)
	mockContent := "line1\nline2\nline3\nline4\nline5\nline6\nline7\nline8\n\nline1!\nline2!\nline3!\nline4!\nline5!\nline6!\nline7!\nline8!"

	charMap := buildCharMap(mockContent)

	// Should have space character (ASCII 32)
	spacePattern, exists := charMap[' ']
	if !exists {
		t.Error("Space character not found")
	}
	if len(spacePattern) != 8 {
		t.Errorf("Space should have 8 lines, got %d", len(spacePattern))
	}
	if spacePattern[0] != "line1" {
		t.Errorf("Expected 'line1', got %q", spacePattern[0])
	}

	// Should have exclamation character (ASCII 33)
	exclPattern, exists := charMap['!']
	if !exists {
		t.Error("Exclamation character not found")
	}
	if len(exclPattern) != 8 {
		t.Errorf("Exclamation should have 8 lines, got %d", len(exclPattern))
	}
	if exclPattern[0] != "line1!" {
		t.Errorf("Expected 'line1!', got %q", exclPattern[0])
	}
}

// Test buildCharMap with empty content
func TestBuildCharMap_EmptyContent(t *testing.T) {
	charMap := buildCharMap("")
	if len(charMap) != 1 {
		t.Errorf("Expected map with 1 character (space), got %d characters", len(charMap))
	}
	// Should have space character with empty content
	spacePattern, exists := charMap[' ']
	if !exists {
		t.Error("Space character should exist")
	}
	if len(spacePattern) != 1 || spacePattern[0] != "" {
		t.Errorf("Expected space to have one empty string, got %v", spacePattern)
	}
}

// Test renderWord with mock character map
func TestRenderWord_SingleCharacter(t *testing.T) {
	// Create mock character map
	charMap := map[rune][]string{
		'A': {"  A  ", " A A ", "AAAAA", "A   A", "A   A", "     ", "     ", "     "},
		' ': {"     ", "     ", "     ", "     ", "     ", "     ", "     ", "     "},
	}

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	renderWord("A", charMap)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("Expected 8 lines, got %d", len(lines))
	}
	if lines[0] != "  A  " {
		t.Errorf("Expected '  A  ', got %q", lines[0])
	}
}

// Test renderWord with multiple characters
func TestRenderWord_MultipleCharacters(t *testing.T) {
	charMap := map[rune][]string{
		'A': {"A", "A", "A", "A", "A", "A", "A", "A"},
		'B': {"B", "B", "B", "B", "B", "B", "B", "B"},
		' ': {" ", " ", " ", " ", " ", " ", " ", " "},
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	renderWord("AB", charMap)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("Expected 8 lines, got %d", len(lines))
	}
	if lines[0] != "AB" {
		t.Errorf("Expected 'AB', got %q", lines[0])
	}
}

// Test renderWord with missing character (should show Invalid!)
func TestRenderWord_MissingCharacter(t *testing.T) {
	charMap := map[rune][]string{
		' ': {"_", "_", "_", "_", "_", "_", "_", "_"},
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	renderWord("X", charMap) // X doesn't exist, should show Invalid!

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if lines[0] != " I " {
		t.Errorf("Expected 'Invalid!' pattern (' I '), got %q", lines[0])
	}
}

// Test renderText with empty string
func TestRenderText_EmptyString(t *testing.T) {
	charMap := map[rune][]string{}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	renderText("", charMap)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if output != "" {
		t.Errorf("Expected no output for empty string, got %q", output)
	}
}

// Test renderText with literal \n
func TestRenderText_LiteralNewline(t *testing.T) {
	charMap := map[rune][]string{
		'A': {"A", "A", "A", "A", "A", "A", "A", "A"},
		'B': {"B", "B", "B", "B", "B", "B", "B", "B"},
		' ': {" ", " ", " ", " ", " ", " ", " ", " "},
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	renderText("A\\nB", charMap) // Literal \n should become real newline

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Should have A (8 lines) + B (8 lines) = 16 lines total
	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) != 16 {
		t.Errorf("Expected 16 lines (8 for A + 8 for B), got %d", len(lines))
	}
}

// Test renderText with empty line
func TestRenderText_EmptyLine(t *testing.T) {
	charMap := map[rune][]string{}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	renderText("\n", charMap) // Just a newline should print one empty line

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if output != "\n" {
		t.Errorf("Expected single newline, got %q", output)
	}
}

// Integration test with real banner file
func TestIntegration_StandardBanner(t *testing.T) {
	content, err := readBannerFile("standard", "assets/")
	if err != nil {
		t.Skipf("Skipping integration test: %v", err)
	}

	charMap := buildCharMap(content)

	// Test that common characters exist and have correct height
	testChars := []rune{' ', 'A', 'a', '0', '!'}
	for _, char := range testChars {
		pattern, exists := charMap[char]
		if !exists {
			t.Errorf("Character %q not found in standard banner", char)
			continue
		}
		if len(pattern) != 8 {
			t.Errorf("Character %q should have 8 lines, got %d", char, len(pattern))
		}
	}
}
