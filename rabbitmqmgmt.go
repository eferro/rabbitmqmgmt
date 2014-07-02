/* greet.go */
package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "rabbitmqmgmt"
	app.Usage = "queue "
	app.Action = func(c *cli.Context) {
		println("testing!")
	}

	app.Commands = []cli.Command{
		{
			Name:      "queue",
			ShortName: "q",
			Usage:     "options for queues",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new queue",
					Action: func(c *cli.Context) {
						println("new queue: ", c.Args().First())
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing queue",
					Action: func(c *cli.Context) {
						println("removed queue: ", c.Args().First())
					},
				},
			},
		},
    {
      Name:      "exchange",
      ShortName: "q",
      Usage:     "options for queues",
      Subcommands: []cli.Command{
        {
          Name:  "add",
          Usage: "add a new exchange",
          Action: func(c *cli.Context) {
            println("new exchange: ", c.Args().First())
          },
        },
        {
          Name:  "remove",
          Usage: "remove an existing exchange",
          Action: func(c *cli.Context) {
            println("removed exchange: ", c.Args().First())
          },
        },
      },
    },
	}
	app.Run(os.Args)
}
