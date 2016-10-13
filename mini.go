package main

import "gopkg.in/gin-gonic/gin.v1"
import "net/http"

type MiniInterface interface {
    HandlePop(*gin.Context)
    HandlePush(*gin.Context)
}

type mini struct {
    Host string
}

type Mini mini

// Handle the request to pop a new item
// from a specific queue
func (m *Mini) HandlePop(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "data":"test",
    })
}

// Handle the request to push a new item
// onto a specific queue
func (m *Mini) HandlePush(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "data":"test",
    })
}