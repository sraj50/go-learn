package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
  // the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Get a greeting message and print it.
	message, err := greetings.Hello("Sid")
	// if an error was returned, print it to the console and exit program
	if err != nil {
		log.Fatal(err)
	}

	// if no error returned, print the returned message to console
	fmt.Println(message)
}