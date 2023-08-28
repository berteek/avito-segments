package main

import (
    "net/http"
    "log"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong!!!",
        })
    })
    err := r.Run()
    if err != nil {
        log.Fatalf("could not run the engine: %v", err)
    }
}
