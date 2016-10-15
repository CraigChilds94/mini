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
    queue := c.Param("queue")

    c.JSON(http.StatusOK, gin.H{
        "queue": queue,
    })
}

// Handle the request to push a new item
// onto a specific queue
func (m *Mini) HandlePush(c *gin.Context) {
    data := c.PostForm("message")
    queue := c.Param("queue")

    c.JSON(http.StatusOK, gin.H{
        "queue": queue,
        "data":  data,
    })
}
