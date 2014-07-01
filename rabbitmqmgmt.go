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

  app.Run(os.Args)
}