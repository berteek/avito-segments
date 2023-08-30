package main

import (
    "log"

    "github.com/berteek/avito-segments/internal/api"
)

func main() {
    e := api.MakeEngine()
    err := e.Run()
    if err != nil {
        log.Fatalf("could not run the engine: %v", err)
    }
}
