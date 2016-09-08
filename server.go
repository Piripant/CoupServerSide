package main

import (
    "fmt"
    "time"
    "net/http"
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

    new_server := server {s_info[0], s_info[1], s_info[2], true}
    server_list = append(server_list, new_server)
}

// Resets a server as online
func reset_server (w http.ResponseWriter, r *http.Request) {
    server_ip := r.URL.Query().Get("ip")
    server_port := r.URL.Query().Get("port")
    
    if server_ip == "" || server_port == "" {
        return
    }
    
    for i := 0; i < len(server_list); i++ {
        if server_list[i].ip == server_ip && server_list[i].port == server_port {
            server_list[i].online = true;
            return
        }
    }
}

func display_servers(w http.ResponseWriter, r *http.Request) {
    for i := 0; i < len(server_list); i++ {
        fmt.Fprintf(w, server_list[i].name + "|" + server_list[i].ip + "|" + server_list[i].port + "\n")
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
                // Sets server as offline, if none will reset it ->
                // will get deleted the next checking
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
    http.HandleFunc("/add", add_server)
    http.HandleFunc("/reset", reset_server)
    
    fmt.Printf("Web pages handlers initializated\n")
    
    go check_servers()
    
    fmt.Printf("Server checker started\n")
    
    http.ListenAndServe(":8080", nil)
}