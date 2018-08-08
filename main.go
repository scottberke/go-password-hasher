package main

import (
  "log"
  "flag"
  "time"
  "github.com/scottberke/password_hasher/server"
)

func main() {
  // Setup flags for command line args
  portPtr := flag.Int("port", 8080, "a port to start the server on")
  delayPtr := flag.Int("delay", 5, "number of seconds to delay hash response")
  flag.Parse()

  log.Printf("main: starting HTTP server on port %d", *portPtr)

  // Setup server and done channel to block main() termination
  // NewServer accepts done channel so endpoint can signal termination
  // is acceptable
  done := make(chan bool)
  server := server.NewServer(*portPtr, time.Duration(*delayPtr), done)

  // Start server in a goroutine. Probably not nessecary since the
  // blocking channel is terminated in the shutdown handler
  go func() {
    if err := server.ListenAndServe(); err != nil {
      log.Printf("Error occurred: %v", err)
    }
  }()

  // Call ShutdownServer but have it's execution blocked
  server.ShutdownServer()

  // Block termination 
  <-done
  log.Printf("Exiting.")
}
