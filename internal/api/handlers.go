package api

import (
    "net/http"
    "context"
    "fmt"
    "errors"

    "github.com/gin-gonic/gin"
)

func MakeCreateUserHandler(s SegmentService) func(c *gin.Context) {
    logError := makeLogErrorFunc("unable to create a user")
    return func(c *gin.Context) {
        err := s.CreateUser(context.Background())
        if err != nil {
            logError(c, err)
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
            "message": "successfully created a user",
        })
    }
}

func MakeCreateSegmentHandler(s SegmentService) func(c *gin.Context) {
    logError := makeLogErrorFunc("unable to create a segment")
    return func(c *gin.Context) {
        slug, err := getSlugFromContext(c)
        if err != nil {
            logError(c, err)
            return
        }

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

func MakeDeleteSegmentHandler(s SegmentService) func(c *gin.Context) {
    logError := makeLogErrorFunc("unable to delete a segment")
    return func(c *gin.Context) {
        slug, err := getSlugFromContext(c)
        if err != nil {
            logError(c, err)
            return
        }

        err = s.DeleteSegment(context.Background(), slug)
        if err != nil {
            logError(c, err)
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
            "message": fmt.Sprintf("successfully deleted a segment: %s", slug),
        })
    }
}

func MakeAddAndDeleteSegmentsFromUserHandler(s SegmentService) func(c *gin.Context) {
    logError := makeLogErrorFunc("unable to add or delete segments from user")
    return func(c *gin.Context) {
        type requestBody struct {
            UserID           int      `json:"user_id"            binding:"required"`
            SegmentsToAdd    []string `json:"segments_to_add"    binding:"required"`
            SegmentsToDelete []string `json:"segments_to_delete" binding:"required"`
        }

        parseRequestBodyFromContext := func(c *gin.Context) (requestBody, error) {
            var rb requestBody
            err := c.ShouldBindJSON(&rb)
            if err != nil {
                return requestBody{}, errors.New("invalid json structure")
            }
            return rb, nil
        }

        reqBody, err := parseRequestBodyFromContext(c)
        if err != nil {
            logError(c, err)
            return
        }

        userID           := reqBody.UserID
        segmentsToAdd    := reqBody.SegmentsToAdd
        segmentsToDelete := reqBody.SegmentsToDelete

        err = s.AddAndDeleteSegmentsFromUser(context.Background(), userID, segmentsToAdd, segmentsToDelete)
        if err != nil {
            logError(c, err)
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
            "message": fmt.Sprintf("successfully added and deleted segments from user (id=%d)", userID),
            "segments_to_add": segmentsToAdd,
            "segments_to_delete": segmentsToDelete,
        })
    }
}

func MakeGetActiveSegmentsForUserHandler(s SegmentService) func(c *gin.Context) {
    logError := makeLogErrorFunc("unable to get active segments for user")
    return func(c *gin.Context) {
        reqBody := struct {
            UserID int `json:"user_id" binding:"required"`
        }{}
        err := c.ShouldBindJSON(&reqBody)
        if err != nil {
            logError(c, errors.New("invalid json structure"))
            return
        }

        userID := reqBody.UserID
        activeSegments, err := s.GetActiveSegmentsForUser(context.Background(), userID)
        if err != nil {
            logError(c, err)
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
            "message": fmt.Sprintf("active segments for user (id=%d) found", userID),
            "segments": activeSegments,
        })
    }
}

func getSlugFromContext(c *gin.Context) (string, error) {
    reqBody := struct {
        Slug string `json:"slug" binding:"required"`
    }{}
    err := c.ShouldBindJSON(&reqBody)
    if err != nil {
        return "", errors.New("invalid json structure")
    }
    return reqBody.Slug, nil
}

func makeLogErrorFunc(message string) func(c *gin.Context, err error) {
    return func(c *gin.Context, err error) {
        c.JSON(http.StatusOK, gin.H{
            "status": "error",
            "message": fmt.Sprintf("%s: %v", message, err),
        })
    }
}
