package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	//Return a greeting that embeds the name in a message.
	var message string
	if name == "" {
		return message, errors.New("empty name")
	}
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
