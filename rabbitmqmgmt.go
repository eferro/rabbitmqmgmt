/* greet.go */
package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func queue_create(amqp_uri string, queue_name string, durable bool, auto_delete bool) {
	conn, err := amqp.Dial(amqp_uri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queue_name,
		durable,
		auto_delete,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

}

func queue_remove(amqp_uri string, queue_name string) {
	conn, err := amqp.Dial(amqp_uri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	_, err = ch.QueueDelete(
		queue_name,
		false,	// ifUnused
		false,  // ifEmpty
		false)   // noWait
	failOnError(err, "Failed to declare a queue")
}

func main() {
	amqp_uri := "amqp://guest:guest@localhost:5672/"

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
						queue_create(amqp_uri, c.Args().First(), c.Bool("durable"), c.Bool("auto-delete"))
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing queue",
					Action: func(c *cli.Context) {
						queue_remove(amqp_uri, c.Args().First())
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
