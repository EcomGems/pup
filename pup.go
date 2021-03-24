package main

import (
	"github.com/urfave/cli/v2"
	"gitlab.com/ecomgems/pup/cmd/serve"
	"gitlab.com/ecomgems/pup/cmd/use"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:                 "pup",
		Usage:                "Tiny file sync through network solution",
		Version:              "1.0.0",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			use.GetCommand(),
			serve.GetCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
