package main

import (
	"fmt"
	"log"
	"playing-with-go/greetings"

	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Orestes", "Rosa", "Irati"}

	// Request a greeting message.
	//message, err := greetings.Hello("Orestes")
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatalf("Fatal error occured : %v", err)
	}

	fmt.Println(messages)

	fmt.Println()
	fmt.Println()
	fmt.Printf("This is a 'Go' quote : %s\n", quote.Go())
	fmt.Printf("This is a 'Glass' quote : %s\n", quote.Glass())
	fmt.Printf("This is a 'Hello' quote : %s\n", quote.Hello())
	fmt.Printf("This is a 'Opt' quote : %s\n", quote.Opt())
}
