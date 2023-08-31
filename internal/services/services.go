package services

import (
    "context"
)

type SegmentRepository interface {
    CreateUser(ctx context.Context) error
    CreateSegment(ctx context.Context, slug string) error
    DeleteSegment(ctx context.Context, slug string) error
    GetActiveSegmentsForUser(ctx context.Context, userID int) ([]string, error)
    AddAndDeleteSegmentsFromUser(ctx context.Context, userID int, segAdd []string, segDel []string) error
}

type SegmentService struct {
    repo SegmentRepository
}

func NewSegmentService(repo SegmentRepository) SegmentService {
    return SegmentService{
        repo: repo,
    }
}

func (s SegmentService) CreateUser(ctx context.Context) error {
    err := s.repo.CreateUser(ctx)
    return err
}

func (s SegmentService) CreateSegment(ctx context.Context, slug string) error {
    err := s.repo.CreateSegment(ctx, slug)
    return err
}

func (s SegmentService) DeleteSegment(ctx context.Context, slug string) error {
    err := s.repo.DeleteSegment(ctx, slug)
    return err
}

func (s SegmentService) AddAndDeleteSegmentsFromUser(ctx context.Context, userID int, segAdd []string, segDel []string) error {
    err := s.repo.AddAndDeleteSegmentsFromUser(context.Background(), userID, segAdd, segDel)
    return err
}

func (s SegmentService) GetActiveSegmentsForUser(ctx context.Context, userID int) ([]string, error) {
    segments, err := s.repo.GetActiveSegmentsForUser(ctx, userID)
    return segments, err
}
