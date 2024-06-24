package main

import (
	"fmt"
	"github.com/swanchain/go-computing-provider/conf"
	"github.com/swanchain/go-computing-provider/internal/computing"
	"github.com/urfave/cli/v2"
	"os"
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

var test2Cmd = &cli.Command{
	Name:  "test2",
	Usage: "test2",
	Action: func(cctx *cli.Context) error {
		fmt.Printf("test2 cmd start!\n")

		cpRepoPath, ok := os.LookupEnv("CP_PATH")
		if !ok {
			return fmt.Errorf("missing CP_PATH env, please set export CP_PATH=<YOUR CP_PATH>")
		}
		if err := conf.InitConfig(cpRepoPath, false); err != nil {
			return fmt.Errorf("load config file failed, error: %+v", err)
		}

		if err := computing.DoTest2(); err != nil {
			return fmt.Errorf("error: %+v", err)
		}

		return nil
	},
}
