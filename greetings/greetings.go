package greetings

import (
	"fmt"
	"errors"
)

// Hello returns greeting for the named person
func Hello(name string) (string, error){
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	
    // Return a greeting that embeds the name in a message.
	// var message string
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}