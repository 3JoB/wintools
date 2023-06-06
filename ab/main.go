package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:   "any bench",
		Usage:  "Benchmark HTTP servers",
		Action: AB,
		Flags:  CliFlags,
	}

	if err := app.Run(os.Args); err != nil {
		PrintUsage()
		panic(err)
	}
}
