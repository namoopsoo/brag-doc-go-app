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
    random_greeting_format := randoFormat()

    message := fmt.Sprintf(random_greeting_format, name )
    //message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {

    some_map := make(map[string]string)

    for _, name := range names {

        some_map[name] = fmt.Sprintf(randoFormat(), name)

    }
    return some_map, nil

}

func randoFormat() string {
    formats := []string {
        "Blah hello %s",
        "MMkay hello %s yea ",
        "Yea %s ok hi.",
    }

    return formats[rand.Intn(len(formats))]

}
