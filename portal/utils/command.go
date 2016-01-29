package utils

import (
	"log"

	"github.com/codegangsta/cli"
)

var ENV = cli.StringFlag{
	Name:   "environment, e",
	Value:  "development",
	Usage:  "Specifies the environment to run this server under (test/development/production).",
	EnvVar: "ITPKG_ENV",
}

func EnvAction(act func(env string, c *cli.Context) error) func(c *cli.Context) {
	return func(c *cli.Context) {
		log.Println("Begin...")
		if e := act(c.String("environment"), c); e == nil {
			log.Println("Done!!!")
		} else {
			log.Fatalln(e)
		}
	}
}
