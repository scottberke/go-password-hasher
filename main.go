package main

import (
  "log"
  "flag"
  "time"
  "github.com/scottberke/password_hasher/server"
)

func main() {
  portPtr := flag.Int("port", 8080, "a port to start the server on")
  delayPtr := flag.Int("delay", 5, "number of seconds to delay hash response")
  flag.Parse()

  log.Printf("main: starting HTTP server on port %d", *portPtr)


  done := make(chan bool)
  server := hashserver.NewServer(*portPtr, time.Duration(*delayPtr), done)

  go func() {
    if err := server.ListenAndServe(); err != nil {
      log.Printf("Error occurred during startup: %v", err)
    }
  }()

  server.ShutdownServer()

  <-done
  log.Printf("Exiting.")
}
