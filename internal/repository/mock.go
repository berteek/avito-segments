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
