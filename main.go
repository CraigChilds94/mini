package main

// Start running the server which
// handles the queue
func main() {
    http := new(Http)
    http.Start(Mini{
        Host: "localhost:8101",
    })
}
