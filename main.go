package main

import (
	"fmt"
	"os"
	"os/exec"
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
					Name:  "pem",
					Usage: "Specify the source of the pem key from `FILE`",
				},
				cli.StringFlag{
					Name:  "ppk",
					Usage: "Specify the destination of where you want to save ppk from `FILE`",
				},
			},
			Action: func(c *cli.Context) error {
				pemKeyName := c.String("pem")
				ppkKeyName := c.String("ppk")
				fmt.Println(pemKeyName)
				fmt.Println(ppkKeyName)
				cmd := exec.Command("puttygen", pemKeyName, "-O", "private", "-o", ppkKeyName)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				cmd.Run()
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}
