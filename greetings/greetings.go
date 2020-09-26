package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	
	// if no name is given, return error with message
	if name == "" {
		return "", errors.New("empty name")
	}

	// create a message using a random format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// set initial values for variables used in the function
func init() {
	rand.Seed(time.Now().UnixNano())
}

// returns one among a set of greetings
// returned message is random
func randomFormat() string {

	// a slice of message formats
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v Well met!",
	}

	// return random message format
	// specifies a random index for slice of formats
	return formats[rand.Intn(len(formats))]

}