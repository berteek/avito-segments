package api

import (
    "context"

    "github.com/gin-gonic/gin"
)

type SegmentService interface {
    CreateSegment(ctx context.Context, slug string) error
    DeleteSegment(ctx context.Context, slug string) error
}

func MakeEngine(s SegmentService) *gin.Engine {
    e := gin.Default()
    e.POST("/create", MakeCreateSegmentHandler(s))
    return e
}
