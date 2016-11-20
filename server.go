package main

import (
    "fmt"
    "time"
    "net/http"
    "os"
)

type server struct {
    name    string
    ip      string
    port    string
    players string
    online  bool
}

var server_list []server

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Seems that the server is up and running!")
}


// Resets a server as online
func reset_server(w http.ResponseWriter, r *http.Request) {
    server_name :=  r.URL.Query().Get("n")
    server_ip := r.URL.Query().Get("i")
    server_port := r.URL.Query().Get("p")
    server_players := r.URL.Query().Get("l")
    
    if server_name == "" || server_ip == "" || server_port == "" || server_players == "" {
        fmt.Printf("Bad refresh request received\n")
        return
    }
    
    for i := 0; i < len(server_list); i++ {
        if server_list[i].ip == server_ip && server_list[i].port == server_port {
            server_list[i].players = server_players
            server_list[i].online = true
            return
        }
    }

    // There was no server with that data
    new_server := server {server_name, server_ip, server_port, server_players, true}
    server_list = append(server_list, new_server)

    fmt.Printf("Server added\n");
}

func display_servers(w http.ResponseWriter, r *http.Request) {
    for i := 0; i < len(server_list); i++ {
        fmt.Fprintf(w, server_list[i].name + "|" + server_list[i].ip + "|" + server_list[i].port + "|" + server_list[i].players + "\n")
    }
}

func check_servers() {
    for {
        i := 0
        for i < len(server_list) {
            // Checks if server is online
            if server_list[i].online == false {
                server_list = append(server_list[:i], server_list[i+1:]...)
            } else {
                // Sets server as offline
                // will get deleted the next checking if it hasn't been reset
                server_list[i].online = false
                i++
            }
        }
        
        // Checks again every 4 seconds
        time.Sleep(4 * time.Second)
    }
}

func main() {
    fmt.Printf("Starting the server, beep beep boop\n")
    
    http.HandleFunc("/coupfps", handler)
    http.HandleFunc("/display", display_servers)
    http.HandleFunc("/reset", reset_server)
    
    fmt.Printf("Web pages handlers initializated\n")
    
    go check_servers()
    
    fmt.Printf("Server checker started\n")
    
    port := os.Getenv("PORT")
    http.ListenAndServe(":" + port, nil)
}