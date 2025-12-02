package main

import (
	"testing"
)

// Test that the banner file can be read and is not empty
func TestReadBannerFile(t *testing.T) {
	content, err := readBannerFile("../../assets/standard.txt")
	if err != nil {
		t.Fatalf("Failed to read banner file: %v", err)
	}
	if len(content) == 0 {
		t.Fatal("Banner file is empty")
	}
}

// Test that characters are parsed correctly into the map
func TestBuildCharMap(t *testing.T) {
	content, err := readBannerFile("../../assets/standard.txt")
	if err != nil {
		t.Fatalf("Failed to read banner file: %v", err)
	}

	charMap := buildCharMap(content)

	// Space character should exist
	if _, exists := charMap[' ']; !exists {
		t.Error("Space character (' ') not found in charMap")
	}

	// 'A' should exist and have exactly 8 lines
	pattern, exists := charMap['A']
	if !exists {
		t.Error("Character 'A' not found in charMap")
	} else if len(pattern) != 8 {
		t.Errorf("Character 'A' should have 8 lines, got %d", len(pattern))
	}
}

// Optional simple table-driven test for a few characters
func TestSomeCharactersExist(t *testing.T) {
	content, err := readBannerFile("../../assets/standard.txt")
	if err != nil {
		t.Fatalf("Failed to read banner file: %v", err)
	}

	charMap := buildCharMap(content)

	tests := []struct {
		char rune
		name string
	}{
		{'A', "capital A"},
		{'B', "capital B"},
		{'0', "digit 0"},
		{'9', "digit 9"},
		{'!', "exclamation mark"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pattern, exists := charMap[tt.char]
			if !exists {
				t.Fatalf("Character %q not found in charMap", tt.char)
			}
			if len(pattern) != 8 {
				t.Fatalf("Character %q should have 8 lines, got %d", tt.char, len(pattern))
			}
		})
	}
}
