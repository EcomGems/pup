package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
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
			{
				Name:    "serve",
				Aliases: []string{"s"},
				Usage:   "Serve files on the remote server to be synced",
				Action: func(c *cli.Context) error {
					fmt.Println(c)
					return nil
				},
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name: "uid",

						Usage:       "User ID to be used for files and folder creation",
						DefaultText: "1000",
					},
					&cli.StringFlag{
						Name:        "gid",
						Usage:       "User Group ID to be used for files and folder creation",
						DefaultText: "1000",
					},
				},
			},
			{
				Name:    "use",
				Aliases: []string{"u"},
				Usage:   "Connects to the server to sync remote folder with the local one",
				Action: func(c *cli.Context) error {
					fmt.Println(c.FlagNames())

					for _, ignoredPath := range c.StringSlice("ignore") {
						fmt.Println("Ignore:", ignoredPath)
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringSliceFlag{
						Name:  "ignore",
						Usage: "Folders to be skipped via sync process",
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
