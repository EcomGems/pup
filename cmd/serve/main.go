package serve

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "Serve files on the remote serve to be synced",
		Action: func(c *cli.Context) error {
			fmt.Println(c)
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
