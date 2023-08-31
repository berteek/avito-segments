package main

import (
    "log"
    "os"
    "context"

    "github.com/berteek/avito-segments/internal/api"
    svc  "github.com/berteek/avito-segments/internal/services"
    rep "github.com/berteek/avito-segments/internal/repository"
    "github.com/jackc/pgx/v5"
)

func main() {
    dbURL := os.Getenv("DATABASE_URL")
    conn, err := pgx.Connect(context.Background(), dbURL)
    if err != nil {
        log.Fatalf("unable to connect to database: %v", err)
    }
    defer conn.Close(context.Background())

    r := rep.NewSegmentRepository(conn)
    s := svc.NewSegmentService(r)
    e := api.MakeEngine(s)
    err = e.Run()
    if err != nil {
        log.Fatalf("could not run the engine: %v", err)
    }
}
