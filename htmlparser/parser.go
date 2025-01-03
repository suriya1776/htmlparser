package htmlparser

import "fmt"

func Test() {
	fmt.Println("Print the test")
}

func Parse(filelocation string) (string, error) {
	// File location verifications
	err := fCheck(filelocation)
	if err != nil {
		return "", err
	}

	return "", nil
}
