package main

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Hello! will returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func main() {
	message, err := Hello("Fon!")
	if err != nil {
		log.Fatal(err)
	}
	log.Info(message)
}
