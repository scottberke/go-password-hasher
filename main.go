package main

import (
  "log"
  "github.com/scottberke/password_hasher/server"
)

func main() {
  log.Printf("main: starting HTTP server")
  server := hashserver.NewServer(0)
  done := make(chan bool)

  go func() {
    if err := server.ListenAndServe(); err != nil {
      log.Fatalf("Error occurred during startup: %v", err)
    }
    done <- true
  }()

  server.ShutdownServer()

  <-done
  log.Printf("Exiting.")
}
