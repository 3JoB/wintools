package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "win-touch",
		Usage: "touch [args]... file...",
		Action: Touch,
		Flags: CliFlags,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
