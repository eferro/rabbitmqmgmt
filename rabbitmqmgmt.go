/* greet.go */
package main

import (
	"github.com/codegangsta/cli"
	"github.com/streadway/amqp"
	"os"
)

func queue_create(amqp_uri string, name string, durable bool, auto_delete bool) {
	println("queue create: ", name, durable, auto_delete)

	_, err := amqp.Dial(amqp_uri)
	if err != nil {
		println("Dial: ", err)
	}


}

func queue_remove(name string) {
	println("queue remove: ", name)
}

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
						queue_create("amqp://guest:guest@localhost:5672/", c.Args().First(), c.Bool("durable"), c.Bool("auto-delete"))
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing queue",
					Action: func(c *cli.Context) {
						queue_remove(c.Args().First())
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
