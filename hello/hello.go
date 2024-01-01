package main


import (
    "fmt"
    "log"

    "example/greetings"
)

func main() {
    // logger 
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    message, err := greetings.Hello("")
    if err != nil {
        log.Fatal(err)
    }
    
    // otherwise,
    fmt.Println(message)
}
