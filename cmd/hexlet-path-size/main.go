package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"code"

	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "human",
				Aliases:     []string{"H"},
				Usage:       "human-readable sizes (auto-select unit)",
				HideDefault: true,
				Local:       true,
			},
			&cli.BoolFlag{
				Name:        "all",
				Aliases:     []string{"a"},
				Usage:       "include hidden files and directories",
				HideDefault: true,
				Local:       true,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			human := cmd.Bool("human")
			all := cmd.Bool("all")
			path := cmd.Args().First()
			if path == ""{
				path = "./"
			}
			size, err := code.GetSize(path, human, all)
			if err != nil {
				return err
			}
			fmt.Println(size)
			return nil
		},
	}
	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
