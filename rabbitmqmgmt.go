/* greet.go */
package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "rabbitmqmgmt"
	app.Usage = "rabbitmq queue/exchage/bindings management"

	app.Commands = []cli.Command{
		{
			Name:      "queue",
			ShortName: "q",
			Usage:     "options for queues",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new queue",
					Flags: []cli.Flag{
						cli.BoolFlag{"durable", "queue survive broker restart"},
						cli.BoolFlag{"auto-delete", "queue is deleted when last consumer unsubscribes"},
					},
					Action: func(c *cli.Context) {
						println("new queue: ", c.Args().First(), "durable", c.Bool("durable"), "auto-delete", c.Bool("auto-delete"))
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
			ShortName: "e",
			Usage:     "options for exchanges",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new exchange",
					Flags: []cli.Flag{
						cli.StringFlag{"type", "direct", "exchange type (direct|fanout|topic|Header)"},
						cli.BoolFlag{"durable", "exchanges survive broker restart"},
						cli.BoolFlag{"auto-delete", "exchange is deleted when all queues have finished using it"},
					},
					Action: func(c *cli.Context) {
						println("new exchange: ", c.Args().First(), c.String("type"), "durable", c.Bool("durable"), "auto-delete", c.Bool("auto-delete"))
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
