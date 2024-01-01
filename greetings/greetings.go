package greetings

import (
    "fmt"
    "errors"
    "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // if no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // Wit a name, greet.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}
