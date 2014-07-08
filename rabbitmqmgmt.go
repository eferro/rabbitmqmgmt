package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/streadway/amqp"
	"log"
	"os"
)

const (
	VERSION = "0.0.1"
	ERROR   = -1
)

func validateArgsNumber(c *cli.Context, argsNumber int, msg string) {
	argc := len(c.Args())
	if argc != argsNumber {
		fmt.Println(msg)
		os.Exit(ERROR)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func queue_create(amqp_uri string, queue_name string, durable bool, auto_delete bool, messageTtl int32) {
	conn, err := amqp.Dial(amqp_uri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	args := amqp.Table{}
	if messageTtl > 0 {
		args["x-message-ttl"] = messageTtl
	}

	_, err = ch.QueueDeclare(
		queue_name,
		durable,
		auto_delete,
		false,
		false,
		args,
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
		false, // ifUnused
		false, // ifEmpty
		false) // noWait
	failOnError(err, "Failed to remove a queue")
}

func queue_bind(amqp_uri string, queue_name string, exchange_name string, routing_key string) {
	conn, err := amqp.Dial(amqp_uri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.QueueBind(
		queue_name,
		routing_key,
		exchange_name,
		false, // noWait
		nil)   // args
	failOnError(err, "Failed to bind the queue to the exchange")
}

func queue_unbind(amqp_uri string, queue_name string, exchange_name string, routing_key string) {
	conn, err := amqp.Dial(amqp_uri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.QueueUnbind(
		queue_name,
		routing_key,
		exchange_name,
		nil) // args
	failOnError(err, "Failed to bind the queue to the exchange")
}

func exchange_create(amqp_uri string, queue_name string, exchange_type string, durable bool, auto_delete bool) {
	conn, err := amqp.Dial(amqp_uri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		queue_name,
		exchange_type,
		durable,
		auto_delete,
		false, // internal
		false, // noWait
		nil,   // args
	)
	failOnError(err, "Failed to declare a exchange")
}

func exchange_remove(amqp_uri string, exchange_name string) {
	conn, err := amqp.Dial(amqp_uri)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDelete(
		exchange_name,
		false, // ifUnused
		false) // noWait
	failOnError(err, "Failed to remove exchange")
}

func main() {
	app := cli.NewApp()
	app.Name = "rabbitmqmgmt"
	app.Usage = "rabbitmq queue/exchage/bindings management"
	app.Author = "Eduardo Ferro Aldama"
	app.Email = "eduardo.ferro.aldama@gmail.com"
	app.Version = VERSION
	app.Flags = []cli.Flag{
		cli.StringFlag{"amqp_uri, u", "amqp://guest:guest@localhost:5672/", "broker url (including vhost)"},
	}

	app.Commands = []cli.Command{
		{
			Name:  "queue_add",
			Usage: "add a new queue",
			Flags: []cli.Flag{
				cli.BoolFlag{"durable", "queue survive broker restart"},
				cli.BoolFlag{"auto-delete", "queue is deleted when last consumer unsubscribes"},
				cli.IntFlag{"x-message-ttl", 0, "per-queue message TTL (ms) (0 for no ttl)"},
			},
			Action: func(c *cli.Context) {
				validateArgsNumber(c, 1, "Usage: queue_add queue_name")
				messageTtl := int32(c.Int("x-message-ttl"))
				queue_create(c.GlobalString("amqp_uri"), c.Args().First(), c.Bool("durable"), c.Bool("auto-delete"), messageTtl)
			},
		},
		{
			Name:  "queue_remove",
			Usage: "remove an existing queue",
			Action: func(c *cli.Context) {
				validateArgsNumber(c, 1, "Usage: queue_remove queue_name")
				queue_remove(c.GlobalString("amqp_uri"), c.Args().First())
			},
		},
		{
			Name:  "queue_bind",
			Usage: "bind a queue to a exchange using a ginven topic/routing key",
			Action: func(c *cli.Context) {
				validateArgsNumber(c, 3, "Usage: queue_bind queue exchange routing_key")
				queue_bind(
					c.GlobalString("amqp_uri"),
					c.Args().Get(0), // queue_name
					c.Args().Get(1), // exchange_name
					c.Args().Get(2), // routing_key/topic
				)
			},
		},
		{
			Name:  "queue_unbind",
			Usage: "remove an existing binding",
			Action: func(c *cli.Context) {
				validateArgsNumber(c, 3, "Usage: queue_unbind queue exchange routing_key")
				queue_unbind(
					c.GlobalString("amqp_uri"),
					c.Args().Get(0), // queue_name
					c.Args().Get(1), // exchange_name
					c.Args().Get(2), // routing_key/topic
				)
			},
		},
		{
			Name:  "exchange_add",
			Usage: "add a new exchange",
			Flags: []cli.Flag{
				cli.StringFlag{"type", "direct", "exchange type (direct|fanout|topic|Header)"},
				cli.BoolFlag{"durable", "exchanges survive broker restart"},
				cli.BoolFlag{"auto-delete", "exchange is deleted when all queues have finished using it"},
			},
			Action: func(c *cli.Context) {
				validateArgsNumber(c, 1, "Usage: exchange_add exchange_name")
				exchange_create(c.GlobalString("amqp_uri"), c.Args().First(), c.String("type"), c.Bool("durable"), c.Bool("auto-delete"))
			},
		},
		{
			Name:  "exchange_remove",
			Usage: "remove an existing exchange",
			Action: func(c *cli.Context) {
				validateArgsNumber(c, 1, "Usage: exchange_remove exchange_name")
				exchange_remove(c.GlobalString("amqp_uri"), c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
