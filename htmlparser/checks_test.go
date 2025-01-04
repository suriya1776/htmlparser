package htmlparser

import (
	"os"
	"strings"
	"testing"
)

// Test function for fCheck
func TestFCheck(t *testing.T) {
	// Test case 1: File does not exist
	t.Run("FileDoesNotExist", func(t *testing.T) {
		err := fCheck("nonexistentfile.html")
		if err == nil || !strings.Contains(err.Error(), "file does not exist") {
			t.Errorf("expected error 'file does not exist', got: %v", err)
		}
	})

	// Test case 2: Path is a directory, not a file
	t.Run("PathIsDirectory", func(t *testing.T) {
		dirPath := "testdir" // Make sure this is an existing directory
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			t.Fatalf("failed to create test directory: %v", err)
		}
		defer os.Remove(dirPath)

		err = fCheck(dirPath)
		if err == nil || !strings.Contains(err.Error(), "provided path is a directory") {
			t.Errorf("expected error 'provided path is a directory', got: %v", err)
		}
	})

	// Test case 3: Invalid file extension
	t.Run("InvalidFileExtension", func(t *testing.T) {
		err := fCheck("testfile.txt")
		if err == nil || !strings.Contains(err.Error(), "invalid file type expected html file") {
			t.Errorf("expected error 'invalid file type expected html file', got: %v", err)
		}
	})

	// Test case 4: File exists and can be opened
	t.Run("FileExistsAndCanBeOpened", func(t *testing.T) {
		// Create a test HTML file
		testFile := "testfile.html"
		content := "<html><body>Test</body></html>"
		err := os.WriteFile(testFile, []byte(content), 0644)
		if err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}
		defer os.Remove(testFile)

		err = fCheck(testFile)
		if err != nil {
			t.Errorf("unexpected error when opening file: %v", err)
		}
	})
}
