// Package main initializes the cli of go-builder
package main

import (
	"log"
	"os"

	cli "github.com/nikhilsbhat/renderer/cli"
)

//This function is responsible for starting the application.
func main() {
	app := cli.CliApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
