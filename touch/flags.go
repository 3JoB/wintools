package main

import "github.com/urfave/cli/v2"

var CliFlags = []cli.Flag{
	CliFlagModifyTime,
}

var CliFlagModifyTime = &cli.BoolFlag {
	Name:  "m",
    Value: false,
}