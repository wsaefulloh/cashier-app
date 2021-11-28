package main

import (
	"log"
	"os"

	"github.com/wsaefulloh/go-solid-principle/configs/command"
)

func main() {
	err := command.Run(os.Args[1:])
	if err != nil {
		log.Fatal("Unable to run app")
	}
}
