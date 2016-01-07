package main

import (
	"log"

	_ "github.com/itpkg/reading/api/assets"
	_ "github.com/itpkg/reading/api/auth"
	"github.com/itpkg/reading/api/core"
	_ "github.com/itpkg/reading/api/site"
	_ "github.com/lib/pq"
)

func main() {
	if err := core.Run(); err != nil {
		log.Fatalln(err)
	}
}
