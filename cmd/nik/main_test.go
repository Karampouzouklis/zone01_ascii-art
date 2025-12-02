package main

import "testing"

func TestReadBannerFile(t *testing.T) {
	content, err := readBannerFile("../../assets/standard.txt")

	if err != nil {
		t.Errorf("Failed to read File: %v", err)
	}

	if len(content) == 0 {
		t.Errorf("File is empty")
	}
}

func TestParseCharacters(t *testing.T) {
	content, _ := readBannerFile("../../assets/standard.txt")
	charMap := parseCharacters(content)

	// Test space character exists
	if _, exists := charMap[' ']; !exists {
		t.Error("Space character not found")
	}

	// Test 'A' character exists and has 8 lines
	if pattern, exists := charMap['A']; exists {
		if len(pattern) != 8 {
			t.Errorf("Character 'A' should have 8 lines, got %d", len(pattern))
		}
	} else {
		t.Error("Character 'A' not found")
	}
}
