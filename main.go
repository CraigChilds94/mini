package main

import os "os"

// Start running the server which
// handles the queue
func main() {
    http := new(Http)
    store := NewStore()

    store.Put([]byte("This is a test message"))

    http.Start(Mini{
        Host:  os.Getenv("HOST"),
        Store: store,
    })
}
