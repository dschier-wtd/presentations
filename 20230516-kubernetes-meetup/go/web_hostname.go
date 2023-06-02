package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
    if err != nil {
        fmt.Fprintln(w, "Error getting hostname:", err)
        return
    }
    fmt.Fprintln(w, "Hostname:", hostname)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}
