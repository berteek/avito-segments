package services

import (
    "context"
)

type SegmentRepository interface {
    CreateSegment(ctx context.Context, slug string) error
    DeleteSegment(ctx context.Context, slug string) error
}

type SegmentService struct {
    repo SegmentRepository
}

func NewSegmentService(repo SegmentRepository) SegmentService {
    return SegmentService{
        repo: repo,
    }
}

func (s SegmentService) CreateSegment(ctx context.Context, slug string) error {
    err := s.repo.CreateSegment(ctx, slug)
    return err
}

func (s SegmentService) DeleteSegment(ctx context.Context, slug string) error {
    err := s.repo.DeleteSegment(ctx, slug)
    return err
}
