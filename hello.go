package main

import (
    "fmt"
    "log"
    "github.com/mtclinton/greetings"
)

func main() {
    // Set properties of the defined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Max", "John", "Tim"}

    // Get a greeting message and print it.
    messages, err := greetings.Hellos(names)
    // If an error was returned, print it to the console and
    // exit the program
    if err != nil {
	    log.Fatal(err)
    }

    // If no error was returned, print the returned messages to console
    fmt.Println(messages)
}
