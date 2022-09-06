package main

import (
	"github.com/adamdubey/toolkit"
)

func main() {
	var tools toolkit.Tools

	tools.CreateDirIfNotExist("./test-dir")
}
