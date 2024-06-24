package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var testCmd = &cli.Command{
	Name:  "test",
	Usage: "test",
	Action: func(cctx *cli.Context) error {
		fmt.Printf("test cmd!")
		return nil
	},
}
