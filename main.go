package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi from a server, hosted by buff!")
}

func main() {
    fmt.Printf("Starting the server, beep beep boop")
    port := os.Getenv("PORT")

    http.HandleFunc("/", handler)
    http.ListenAndServe(":" + port, nil)
}