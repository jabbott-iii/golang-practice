package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {

	var name string

	fmt.Print("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Error(err)
		return
	}

}
