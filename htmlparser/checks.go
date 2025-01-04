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
