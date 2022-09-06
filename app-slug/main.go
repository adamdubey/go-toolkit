package main

import (
	"log"

	"github.com/adamdubey/go-toolkit"
)

func main() {
	toSlug := "hEllO WoRld=42?!"

	var tools toolkit.Tools

	slugified, err := tools.Slugify(toSlug)
	if err != nil {
		log.Println(err)
	}

	log.Println(slugified)
}
