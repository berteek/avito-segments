package repository

import (
    "log"
    "context"
)

type MockSegmentRepository struct {}

func NewMockSegmentRepository() MockSegmentRepository {
    return MockSegmentRepository{}
}

func (r MockSegmentRepository) CreateSegment(ctx context.Context, slug string) error {
    log.Printf("[repository] new segment created: %s", slug)
    return nil
}

func (r MockSegmentRepository) DeleteSegment(ctx context.Context, slug string) error {
    log.Printf("[repository] segment deleted: %s", slug)
    return nil
}

func (r MockSegmentRepository) GetActiveSegmentsForUser(ctx context.Context, userID int) ([]string, error) {
    log.Printf("[repository] looking for active segments for user (id=%d)", userID)
    segments := make([]string, 0)
    return segments, nil
}

func (r MockSegmentRepository) AddAndDeleteSegmentsFromUser(ctx context.Context, userID int, segAdd []string, segDel []string) error {
    log.Printf("[repository] adding and deleting segments for user (id=%d): add=%v, delete=%v", userID, segAdd, segDel)
    return nil
}
