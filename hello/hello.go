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

    some_names := []string{"Jobe", "Jorry", "Jarry"}

    greeting_maps, err := greetings.Hellos(some_names)

    //message, err := greetings.Hello("Jobe")
    if err != nil {
        log.Fatal(err)
    }

    for k, v := range greeting_maps{

        fmt.Println(fmt.Sprintf("hmm, %s, %s", k, v))
    }

    
    // otherwise,
}
