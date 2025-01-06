package htmlparser

import (
	"encoding/json"
	"encoding/xml"
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

type HTML struct {
	XMLName xml.Name `xml:"html"`
	Lang    string   `xml:"lang,attr"`
	Head    Head     `xml:"head"`
	Body    Body     `xml:"body"`
}

type Head struct {
	Meta  []Meta `xml:"meta"`
	Title string `xml:"title"`
}

type Meta struct {
	HTTPEquiv string `xml:"http-equiv,attr,omitempty"`
	Content   string `xml:"content,attr,omitempty"`
	Name      string `xml:"name,attr,omitempty"`
}

type Body struct {
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

	// Preprocess to ensure proper XML self-closing tags
	contentStr := string(content)
	contentStr = strings.ReplaceAll(contentStr, "<meta ", "<meta ")
	contentStr = strings.ReplaceAll(contentStr, ">", " />")
	content = []byte(contentStr)

	var data HTML

	err = xml.Unmarshal([]byte(content), &data)
	if err != nil {
		panic(err)
	}

	jsonOutput, errjson := json.MarshalIndent(data, "", "    ")
	if errjson != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return "", err
	}

	fmt.Println("json:", string(jsonOutput))

	return "", nil
}
