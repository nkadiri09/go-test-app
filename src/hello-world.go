package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    message := os.Getenv("MESSAGE")
    test := os.Getenv("TEST")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        concatenatedMessage := message + " " + test

        // Write the concatenated message to the response
        fmt.Fprintf(w, concatenatedMessage)
    })

    fmt.Println("Starting server on port 8080.")
    http.ListenAndServe(":8080", nil)
}
