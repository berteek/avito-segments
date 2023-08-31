package api

import (
    "context"

    "github.com/gin-gonic/gin"
)

type SegmentService interface {
    CreateSegment(ctx context.Context, slug string) error
    DeleteSegment(ctx context.Context, slug string) error
    GetActiveSegmentsForUser(ctx context.Context, userID int) ([]string, error)
    AddAndDeleteSegmentsFromUser(ctx context.Context, userID int, segAdd []string, segDel []string) error
}

func MakeEngine(s SegmentService) *gin.Engine {
    e := gin.Default()
    e.POST("/create_seg", MakeCreateSegmentHandler(s))
    e.POST("/delete_seg", MakeDeleteSegmentHandler(s))
    e.POST("/add_del_seg_from_user", MakeAddAndDeleteSegmentsFromUserHandler(s))
    e.POST("/active_seg_for_user", MakeGetActiveSegmentsForUserHandler(s))
    return e
}
