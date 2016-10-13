package main

import "gopkg.in/gin-gonic/gin.v1"

type HttpInterface interface {
    Start(Mini)
}

type h struct {}
type Http h

func (h *Http) Start(mini Mini) {
    r := gin.Default()
    
    r.GET("/pop", func(c *gin.Context) {
        mini.HandlePop(c)
    })

    r.Run(mini.Host)
}