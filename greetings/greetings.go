package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns greeting for the named person
func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	// var message string
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// randomFormat returns one of a set of greeting messages. The returned
	// message is selected at random.
	return formats[rand.Intn(len(formats))]
}
