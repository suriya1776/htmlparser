package htmlparser

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Test() {
	fmt.Println("Print the test")
}

func Parse(filelocation string) (string, error) {
	// File location verifications
	err := fCheck(filelocation)
	if err != nil {
		return "", err
	}

	content, err := parse(filelocation)
	if err != nil {
		return "", err
	}
	return content, nil
}

func parse(filelocation string) (string, error) {

	file, err := os.Open(filelocation)
	if err != nil {
		return "", fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("unable to read file: %v", err)
	}

	// Remove all spaces, tabs, and newlines from the string
	processedContent := strings.Join(strings.Fields(string(content)), "")

	// Validate the HTML structure
	if err := validateHTML(processedContent); err != nil {
		return "", fmt.Errorf("invalid HTML: %v", err)
	}

	return processedContent, nil
}
