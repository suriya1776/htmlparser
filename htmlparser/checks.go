package htmlparser

import (
	"fmt"
	"os"
	"regexp"
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

func validateHTML(htmlContent string) error {
	// Regex to match opening, closing, and self-closing HTML tags
	tagRegex := regexp.MustCompile(`</?([a-zA-Z0-9]+)(\s+[^>]*)?>`)

	// Regex to match <script> tags (to ignore their content)
	scriptRegex := regexp.MustCompile(`(?i)<script[^>]*>.*?</script>`)

	// Remove <script> content from validation
	htmlWithoutScripts := scriptRegex.ReplaceAllString(htmlContent, "")

	// Find all tags in the remaining content
	matches := tagRegex.FindAllStringSubmatch(htmlWithoutScripts, -1)

	var stack []string
	for _, match := range matches {
		tag := strings.ToLower(match[1]) // Convert tag name to lowercase for case-insensitive matching
		fullTag := match[0]

		// Skip self-closing tags
		if strings.HasSuffix(fullTag, "/>") {
			continue
		}

		if strings.HasPrefix(fullTag, "</") {
			// Closing tag: check if it matches the top of the stack
			if len(stack) == 0 || stack[len(stack)-1] != tag {
				return fmt.Errorf("unmatched closing tag: </%s>", tag)
			}
			stack = stack[:len(stack)-1] // Pop the matched tag
		} else {
			// Opening tag: push onto the stack
			stack = append(stack, tag)
		}
	}

	// If stack is not empty, there are unmatched opening tags
	if len(stack) > 0 {
		return fmt.Errorf("unmatched opening tag: <%s>", stack[len(stack)-1])
	}

	return nil
}
