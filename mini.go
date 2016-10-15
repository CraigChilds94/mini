package main

import "gopkg.in/gin-gonic/gin.v1"
import "net/http"

type MiniInterface interface {
    HandlePop(*gin.Context)
    HandlePush(*gin.Context)
}

type mini struct {
    Host  string
    Store *Store
}

type Mini mini

// Handle the request to pop a new item
// from a specific queue
func (m *Mini) HandlePop(c *gin.Context) {
    queue := c.Param("queue")

    if queue == "" {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Queue with name not found",
        })
    }

    data, err := m.Store.Next(queue)

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Unable to retrieve item from the queue",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "queue": queue,
        "data":  data,
    })
}

// Handle the request to push a new item
// onto a specific queue
func (m *Mini) HandlePush(c *gin.Context) {
    data := c.PostForm("message")
    queue := c.Param("queue")

    err := m.Store.Put(queue, []byte(data))

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Unable to put item onto the queue",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "queue":  queue,
        "pushed": true,
    })
}
