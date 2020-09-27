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
	// message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprintf(randomFormat())
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message
func Hellos(names []string) (map[string]string, error) {

	// map associated names with message
	messages := make(map[string]string)

	// loop through slice of names, call Hello function to get greeting for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// associate greeting with name
		messages[name] = message
	}
	return messages, nil

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