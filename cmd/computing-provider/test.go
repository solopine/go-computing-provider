package main

import (
	"fmt"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/urfave/cli/v2"
)

var testCmd = &cli.Command{
	Name:  "test",
	Usage: "test",
	Action: func(cctx *cli.Context) error {
		fmt.Printf("test cmd start!\n")

		if err := computing.DoTest(); err != nil {
			return fmt.Errorf("error: %+v", err)
		}

		return nil
	},
}
