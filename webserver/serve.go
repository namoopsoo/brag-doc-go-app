
package main

import (
    "fmt"
    "net/http"
    // "time"
)

func worker(id int, jobs <-chan http.Request) {
    for req := range jobs {
        fmt.Printf("Worker %d processing request: %v\n", id, req)
        // Process the request...
    }
}

func main() {
    jobs := make(chan http.Request, 100) // Buffered channel

    // Start worker goroutines
    for w := 1; w <= 3; w++ {
        go worker(w, jobs)
    }

    // HTTP listener (simplified example)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Received request:", r)
        jobs <- *r // Send the request to the channel
        fmt.Fprintln(w, "Request is being processed")
    })

    http.ListenAndServe(":8080", nil)
}

