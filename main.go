package main

import (
	"os"
	"sort"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "PEMtoPPK"
	app.Usage = "convert pem to ppk"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Sudhir Ganesan",
			Email: "ganesansudhir18@gmail.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "convert",
			Aliases: []string{"c"},
			Usage:   "convert a pem file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "key, k",
					Usage:  "Specify name of pem key from `FILE`",
					EnvVar: "KEY_NAME",
				},
			},
			Action: func(c *cli.Context) error {
				//Code goes here
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}
