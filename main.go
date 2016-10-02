package main

import (
    server "./server"
    fmt "fmt"
    http "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    store := server.NewStore()

    err := store.Put([]byte("Some data"))

    data, err := store.Next()

    if err != nil {
        fmt.Println(err)
    }

    fmt.Fprintf(w, "%s", string(data))
}

// Start running the server which
// handles the queue
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8090", nil)
}
