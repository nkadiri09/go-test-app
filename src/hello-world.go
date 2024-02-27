package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    message := os.Getenv("MESSAGE")
    testmessage := os.Getenv("TEST")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, message)
        fmt.Fprintf(w, message)
    })

    fmt.Println("Starting server on port 8080.")
    http.ListenAndServe(":8080", nil)
}
