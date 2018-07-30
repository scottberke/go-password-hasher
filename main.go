package main

import (
  "log"
  "github.com/scottberke/password_hasher/hash_server"
)


func main() {
  if err := hashserver.HashServer(); err != nil {
    log.Fatalf("Error occurred during startup: %v", err)
  }

  log.Print("Exiting.")
}
