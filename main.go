package main

import (
    "fmt"
    "net/http"
    "os"
)

type server struct {
    name   string
    ip     string
    port   string
    online bool
}

var server_list []server

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi from a server, hosted by buff!")
}

func add_server(w http.ResponseWriter, r *http.Request) {
    fields := []string{"name", "ip", "port"}
    var s_info [3]string

    for i := 0; i < len(fields); i++ {
        s_info[i] = r.URL.Query().Get(fields[i])

        if s_info[i] == "" {
            // The field is missing
            fmt.Printf("Bad add request received")
            return
        }
    }

    //new_server := server {s_info[0], s_info[1], s_info[2], true}
}

func display_servers() {

}

func reset_time_life() {

}

func main() {
    fmt.Printf("Starting the server, beep beep boop")
    port := os.Getenv("PORT")

    http.HandleFunc("/coupfps", handler)
    http.ListenAndServe(":" + port, nil)
}