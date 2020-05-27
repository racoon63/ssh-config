package main

import (
	"fmt"
	"log"
	"os"

	"github.com/racoon63/ssh-config/pkg/commands"
	"github.com/urfave/cli"
)

func main() {
	var name string
	var host string
	var user string
	var port string
	var keyPath string

	app := cli.App{
		Name:            "ssh-config",
		Usage:           "edit your ssh config file",
		UsageText:       "ssh-config [OPTIONS] COMMAND",
		Version:         "0.0.1",
		HideHelpCommand: true,
	}
	app.Commands = []*cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "Prints the current config in ~/.ssh/config",
			Action:  commands.List,
		},
		{
			Name:      "add",
			Usage:     "add a host entry to the ssh config",
			UsageText: "ssh-config add NAME [OPTIONS]",
			Flags: []cli.Flag{
				// &cli.StringFlag{
				// 	Name:        "name",
				// 	Aliases:     []string{"n"},
				// 	Usage:       "Name of the host under which it shall be listed",
				// 	Required:    false,
				// 	Destination: &name,
				// },
				&cli.StringFlag{
					Name:        "host",
					Aliases:     []string{"t"},
					Usage:       "IP or DNS name of target host",
					Required:    true,
					Destination: &host,
				},
				&cli.StringFlag{
					Name:        "user",
					Aliases:     []string{"u"},
					Usage:       "Login username",
					Required:    false,
					Destination: &user,
				},
				&cli.StringFlag{
					Name:        "port",
					Aliases:     []string{"p"},
					Usage:       "Port under which the target host can be reached",
					Required:    false,
					Destination: &port,
				},
				&cli.StringFlag{
					Name:        "key",
					Aliases:     []string{"i"},
					Usage:       "Path to the ssh key to use for the target host.",
					Required:    false,
					Destination: &keyPath,
				},
				&cli.StringFlag{
					Required:    true,
					Destination: &name,
				},
			},
			Action: func(c *cli.Context) error {
				// fmt.Println(c.Args().Get(1))
				// fmt.Println(c.Args().Tail())
				// name = c.Args().Get(0)
				fmt.Println(c.Args())
				fmt.Println(name)
				fmt.Println(host)
				fmt.Println(user)
				fmt.Println(port)
				fmt.Println(keyPath)
				return commands.Add(name, host, user, port, keyPath)
			},
		},
		// {
		// 	Name:  "remove",
		// 	Usage: "remove an existing template",
		// 	Action: func(c *cli.Context) error {
		// 		return nil
		// 	},
		// },
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Prints the ssh-config tool version",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
