package htmlparser

import (
	"fmt"
	"os"
	"strings"
)

func fCheck(filelocation string) error {
	fileinfo, err := os.Stat(filelocation)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %v", err)
		}
		return fmt.Errorf("error accessing the file: %v", err)
	}

	if fileinfo.IsDir() {
		return fmt.Errorf("provided path is a directory")
	}

	if !strings.HasSuffix(filelocation, ".html") {
		return fmt.Errorf("invalid file type expected html file")
	}

	return nil
}

// validateHTML checks for unclosed or mismatched tags
func validateHTML(content string) error {
	stack := []string{} // To track open tags

	for len(content) > 0 {
		// Check for opening or closing tags
		if strings.HasPrefix(content, "<") {
			// Find the next '>' to close the tag
			closeIdx := strings.Index(content, ">")
			if closeIdx == -1 {
				return fmt.Errorf("malformed tag: no closing '>' found")
			}

			tag := content[1:closeIdx]
			content = content[closeIdx+1:]

			// Check for closing tags
			if strings.HasPrefix(tag, "/") {
				tagName := tag[1:]
				if len(stack) == 0 || stack[len(stack)-1] != tagName {
					return fmt.Errorf("mismatched closing tag: </%s>", tagName)
				}
				// Pop the matching opening tag
				stack = stack[:len(stack)-1]
			} else {
				// Handle self-closing tags like <img/>
				if strings.HasSuffix(tag, "/") {
					continue
				}

				// Extract attributes (if any) and strip them for validation
				spaceIdx := strings.IndexAny(tag, " \t")
				if spaceIdx != -1 {
					tag = tag[:spaceIdx] // Keep only the tag name
				}

				// Push the tag onto the stack
				stack = append(stack, tag)
			}
		} else {
			// Move to the next tag (skip plain text)
			nextTagIdx := strings.Index(content, "<")
			if nextTagIdx == -1 {
				break
			}
			content = content[nextTagIdx:]
		}
	}

	// If stack is not empty, some tags were not closed
	if len(stack) > 0 {
		return fmt.Errorf("unclosed tag(s): %v", stack)
	}

	return nil
}
