package main

import (
	"log"

	_ "github.com/itpkg/reading/api/auth"
	_ "github.com/itpkg/reading/api/cms"
	"github.com/itpkg/reading/api/core"
	_ "github.com/itpkg/reading/api/site"
	_ "github.com/lib/pq"
)

func main() {
	if err := core.Run(); err != nil {
		log.Fatalln(err)
	}
}
