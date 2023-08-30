package api

import (
    "github.com/gin-gonic/gin"
)

func MakeEngine() *gin.Engine {
    e := gin.Default()
    e.GET("/ping", SimpleHandler)
    return e
}
