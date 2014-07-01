/* greet.go */
package main

import (
  "os"
  "github.com/codegangsta/cli"
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
        Usage:     "---",
        Action: func(c *cli.Context) {
          println("queue: ", c.Args().First())
        },
    },
    {
        Name:      "exchange",
        ShortName: "e",
        Usage:     "---",
        Action: func(c *cli.Context) {
          println("exchange: ", c.Args().First())
        },
    },
  }

  app.Run(os.Args)
}