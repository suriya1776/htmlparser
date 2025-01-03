package main

import (
	"fmt"

	"github.com/suriya1776/htmlparser/htmlparser"
)

func main() {

	a, err := htmlparser.Parse("testfiles")

	if err != nil {
		fmt.Printf("error exist: %v\n", err)
	} else {
		fmt.Println(a)
	}
}
