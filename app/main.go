package main

import (
	"fmt"

	"github.com/adamdubey/toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(10)
	fmt.Println("Random string:", s)
}
