package htmlparser

import (
	"fmt"
	"io"
	"os"
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

	body, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("unable to read file: %v", err)
	}

	return string(body), nil
}
