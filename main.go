package main

import os "os"

// Start running the server which
// handles the queue
func main() {
    http := new(Http)
    http.Start(Mini{
        Host: os.Getenv("HOST"),
    })
}
