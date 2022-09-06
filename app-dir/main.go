package main

import (
	"github.com/adamdubey/go-toolkit"
)

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExist("./test-dir")
}
