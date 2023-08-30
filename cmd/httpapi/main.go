package main

import (
    "log"

    "github.com/berteek/avito-segments/internal/api"
    svc  "github.com/berteek/avito-segments/internal/services"
    rep "github.com/berteek/avito-segments/internal/repository"
)

func main() {
    r := rep.NewMockSegmentRepository()
    s := svc.NewSegmentService(r)
    e := api.MakeEngine(s)
    err := e.Run()
    if err != nil {
        log.Fatalf("could not run the engine: %v", err)
    }
}
