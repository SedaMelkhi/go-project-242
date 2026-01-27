package main

import (
	"os"
	"context"
	"log"
	"fmt"

	"code"

	"github.com/urfave/cli/v3"
)

func main() {
	size, err := code.GetSize("./")
	if err == nil {
		fmt.Println(size)
	}
	app := &cli.Command{
		Name: "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return nil
		},
	}
	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
	
}
