# Task 7: Basic Tests

## Goal
Learn Go testing basics and verify your code works correctly.

## What You'll Learn
- Go testing framework
- Table-driven tests
- Test file naming conventions

## Steps
1. Create `test/main_test.go`:
   ```go
   package main

   import "testing"

   func TestReadBannerFile(t *testing.T) {
       content, err := readBannerFile("../assets/standard.txt")
       if err != nil {
           t.Errorf("Failed to read banner file: %v", err)
       }
       if len(content) == 0 {
           t.Error("Banner file is empty")
       }
   }
   ```

2. Test character parsing:
   ```go
   func TestParseCharacters(t *testing.T) {
       content, _ := readBannerFile("standard.txt")
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
   ```

3. Run tests:
   ```bash
   go test
   go test -v  # verbose output
   ```

## Success Criteria
- All tests pass
- Understand basic test structure
- Can add new tests for your functions

## Next: Task 8 - Clean up and organize