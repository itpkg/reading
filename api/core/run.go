package core

import (
	"os"

	"github.com/codegangsta/cli"
)

func Run() error {
	app := cli.NewApp()
	app.Name = "reading"
	app.Usage = "I love reading"
	app.Version = "v20160105"
	app.Commands = make([]cli.Command, 0)

	for _, en := range engines {
		cs := en.Shell()
		app.Commands = append(app.Commands, cs...)
	}

	return app.Run(os.Args)
}
