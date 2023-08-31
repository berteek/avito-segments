package repository

import (
    "context"
    "fmt"

    "github.com/jackc/pgx/v5"
    q "github.com/berteek/avito-segments/internal/repository/queries"
)

type SegmentRepository struct {
    conn *pgx.Conn
}

func NewSegmentRepository(conn *pgx.Conn) SegmentRepository {
    return SegmentRepository{
        conn: conn,
    }
}

func (r SegmentRepository) CreateUser(ctx context.Context) error {
    _, err := r.conn.Exec(ctx, q.InsertIntoUsers())
    return err
}

func (r SegmentRepository) CreateSegment(ctx context.Context, slug string) error {
    _, err := r.conn.Exec(ctx, q.InsertIntoSegments(slug))
    return err
}

func (r SegmentRepository) DeleteSegment(ctx context.Context, slug string) error {
    _, err := r.conn.Exec(ctx, q.DeleteFromSegments(slug))
    return err
}

func (r SegmentRepository) GetActiveSegmentsForUser(ctx context.Context, userID int) ([]string, error) {
    type segName struct {
        Name string `db:"name"`
    }
    segments := make([]string, 0)
    rows, err := r.conn.Query(ctx, q.ActiveSegmentsForUser(userID))
    if err != nil {
        return segments, fmt.Errorf("error when getting rows: %v", err)
    }
    segNames, err := pgx.CollectRows(rows, pgx.RowToStructByName[segName])
    if err != nil {
        return segments, fmt.Errorf("error collecting rows: %v", err)
    }
    segments = make([]string, len(segNames))
    for i, sn := range segNames {
        segments[i] = sn.Name
    }
    return segments, nil
}

func (r SegmentRepository) AddAndDeleteSegmentsFromUser(ctx context.Context, userID int, segAdd, segDel []string) error {
    _, err := r.conn.Exec(ctx, q.AddAndDeleteSegmentsFromUser(userID, segAdd, segDel))
    return err
}
