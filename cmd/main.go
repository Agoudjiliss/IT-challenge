package main

import (
    "log"

    "github.com/Agoudjiliss/IT-challenge/internal/server"
)

func main() {
    if err := server.Run(); err != nil {
        log.Fatalf("could not start the server: %v", err)
    }
}
