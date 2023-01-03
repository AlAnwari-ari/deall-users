package main

import (
	"log"

	"github.com/deall-users/cmd"
)

func main() {
	app, err := cmd.App()
	if err != nil {
		log.Fatal(err)
	}

	app.Run()
}
