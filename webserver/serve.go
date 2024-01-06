
package main

import (
    "fmt"
    "net/http"
    // "time"
)


type Packet struct {
    request http.Request
    response_writer http.ResponseWriter
    done chan struct{}
}

func worker(id int, jobs <-chan Packet) {
    for packet := range jobs {
        req := packet.request
        response_writer := packet.response_writer
        fmt.Printf("Worker %d processing request: %v\n", id, req)


        // Process the request...
        fmt.Fprintln(response_writer, "Request is being processed")

        // Signal we are done 
        packet.done <- struct{}{}
    }
}

func main() {
    fmt.Println("Server has started.")
    jobs := make(chan Packet, 100) // Buffered channel

    // Start worker goroutines
    for w := 1; w <= 3; w++ {
        go worker(w, jobs)
    }

    // HTTP listener (simplified example)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Received request:", r)
        done := make(chan struct{})
        packet := Packet{
            request: *r, response_writer: w, done: done,
            }
        jobs <- packet // Send the request to the channel
        <- done
    })

    http.ListenAndServe(":8080", nil)
}

