package htmlparser

import (
	"encoding/xml"
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

type HTML struct {
	XMLName xml.Name `xml:"html"`
	Lang    string   `xml:"lang,attr"`
	Head    Head     `xml:"head"`
	Body    Body     `xml:"body"`
}

type Head struct {
	Title string `xml:"title"`
	Meta  string `xml:"meta"`
}

type Body struct {
	Heading string `xml:"h1"`
	Para    string `xml:"p"`
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

	var data HTML

	err = xml.Unmarshal([]byte(content), &data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Title:", data.Lang)

	return "", nil
}
