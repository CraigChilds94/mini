package main

import "gopkg.in/gin-gonic/gin.v1"

type HttpInterface interface {
    Start(Mini)
}

type h struct {}
type Http h

// Start out http server and bind
// our routes. Add our handlers to each route
func (h *Http) Start(mini Mini) {
    r := gin.Default()
    
    r.GET("/pop", mini.HandlePop)

    r.Run(mini.Host)
}