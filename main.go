package main

import (
    "log"
    "net/http"
)

const (
    port     = ":8080"          // Port the server listens on
    fileDir  = "./shared-files" // Directory to serve files from
)

func main() {
    fs := http.FileServer(http.Dir(fileDir))
    
    http.Handle("/", fs)

    log.Printf("Serving %s on HTTP port %s\n", fileDir, port)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal(err)
    }
}
