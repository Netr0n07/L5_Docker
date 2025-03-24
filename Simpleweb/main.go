package main

import (
    "fmt"
    "net"
    "net/http"
    "os"
)

var version = "default"

func handler(w http.ResponseWriter, r *http.Request) {
    hostname, _ := os.Hostname()
    addrs, _ := net.InterfaceAddrs()

    fmt.Fprintf(w, "<h1>Hostname: %s</h1>", hostname)
    fmt.Fprintf(w, "<h2>IP Addresses:</h2><ul>")
    for _, addr := range addrs {
        fmt.Fprintf(w, "<li>%s</li>", addr.String())
    }
    fmt.Fprintf(w, "</ul>")
    fmt.Fprintf(w, "<p>App Version: %s</p>", version)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
