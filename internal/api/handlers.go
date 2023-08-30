package api

import (
    "net/http"
    "context"
    "fmt"

    "github.com/gin-gonic/gin"
)

type createSegmentRequestBody struct {
    Slug string `json:"slug"`
}

func MakeCreateSegmentHandler(s SegmentService) func(c *gin.Context) {
    logError := func(c *gin.Context, err error) {
        c.JSON(http.StatusOK, gin.H{
            "status": "error",
            "message": fmt.Sprintf("unable to create a segment: %v", err),
        })
    }

    return func(c *gin.Context) {
        var err error

        var reqBody createSegmentRequestBody
        err = c.BindJSON(&reqBody)
        if err != nil {
            logError(c, err)
            return
        }

        slug := reqBody.Slug
        err = s.CreateSegment(context.Background(), slug)
        if err != nil {
            logError(c, err)
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
            "message": fmt.Sprintf("successfully created a segment: %s", slug),
        })
    }
}
