package use

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:    "use",
		Aliases: []string{"u"},
		Usage:   "Connects to the serve to sync remote folder with the local one",
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
	}
}
