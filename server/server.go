package server

import (
  "fmt"
  "log"
  "strings"
  "net/http"
  "time"
  "os"
  "os/signal"
  "syscall"
  "context"
  "encoding/json"
  "github.com/scottberke/password_hasher/hash"
  "github.com/scottberke/password_hasher/timer"
)

// Struct to hold our server and relevant attributes
type Server struct {
  http.Server

  // Channel used to broadcast a shutdown request when recieved via endpoint
  shutdown         chan bool

  // Channel connecting server to scope server is being run in so
  // that the encapsalating scope can block but be unblocked when server
  // shuts down in server package
  doneChan         chan<- bool

  // Counter field to track how many requests the server has received
  requestCount     int

  // Accumulator for duration of requests processed. This allows us to take
  // totalTime / requestCount to get the average request time
  totalTime        time.Duration

  // Requirements specify that the response for a hash should be delayed 5
  // seconds. Adding an attribute allows use to change that when testing etc.
  resDelaySeconds  time.Duration
}

func NewServer(serverPort int, resDelaySeconds time.Duration, done chan<- bool) *Server {
    // Take port supplied via command line arg and build corresponding string
    // for the server to use
    var port strings.Builder
    fmt.Fprintf(&port, ":%d", serverPort)

    // Build our server
    mux := http.NewServeMux()
    server := &Server{
      Server:           http.Server{ Addr: port.String(), Handler: mux },
      shutdown:         make(chan bool),
      resDelaySeconds:  resDelaySeconds,
      doneChan:         done,
    }

    // Setup our handlers
    mux.HandleFunc("/hash", server.hashHandler)
    mux.HandleFunc("/shutdown", server.shutdownHandler)
    mux.HandleFunc("/stats", server.statsHandler)

    return server
}

func (server *Server) hashHandler(w http.ResponseWriter, r *http.Request) {
  // Increment server's requestCount
  server.requestCount += 1
  // Only accept POST calls to this endpoint.
  if r.Method != http.MethodPost {
    w.WriteHeader(405)
    w.Write([]byte(`{"Message": "Method Not Allowed"}`))
  } else {
    // Set start time and defer time tracking until response is served.
    start := time.Now()
    defer timer.AddToTotalTime(start, &server.totalTime)

    log.Printf("Server: Hashing Password")
    base64encodedSha512hash := hashAndEncodePassword(r.FormValue("password"))

    // Keep socket open for desired delay
    time.Sleep(server.resDelaySeconds*time.Second)

    // Write password as string response
    fmt.Fprintf(w, "%s", base64encodedSha512hash)
  }
}

func hashAndEncodePassword(password string) string {
  // Use our hasher and return the base64 encoded SHA512 hash of the password
  hasher := hashencode.Sha512HasherImp
  base64encodedSha512hash := hasher.Hash([]byte(password))

  return base64encodedSha512hash
}

func (server *Server) shutdownHandler(w http.ResponseWriter, r *http.Request) {
  log.Printf("Shutting down server.")
  w.Write([]byte(`{"Message": "Shutdown in progress. Requests Finishing"}`))

  go func() {
    // Unblock server shutdown so its actually called
    server.shutdown <- true
  }()

}

func (server *Server) statsHandler(w http.ResponseWriter, r *http.Request) {
  // Only accept GET requests to this endpoint
  if r.Method != http.MethodGet {
    w.WriteHeader(405)
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(`{"Message": "Method Not Allowed"}`))
  } else {
    // Build our JSON response
    res := make(map[string]int)
    if server.requestCount == 0 {
      res["total"] = 0
      res["average"] = 0
    } else {
      res["total"] = server.requestCount
      res["average"] = int(int64(server.totalTime / time.Microsecond) / int64(server.requestCount))
    }

    // Turn our map into a marshalled JSON response
    data, _ := json.Marshal(res)
    w.Header().Set("Content-Type", "application/json")
    w.Write(data)
  }
}

func (server *Server) ShutdownServer() {
  // Create a channel to listen to system interupts so this
  // can respond to both the shutdown endpoint and ctrl+c
  stop := make(chan os.Signal, 1)
  signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

  // Block until we receive a shutdown signal to one of our channels
  select {
  case sig := <-stop:
     log.Printf("Shutdown request (signal: %v)", sig)
  case sig := <-server.shutdown:
     log.Printf("Shutdown request (/shutdown %v)", sig)
  }

  // Shutdown our server, gracefully waiting for requests to finish,
  // then close the blocking done channel in main() so the app terminates
  go func() {
    if err := server.Shutdown(context.Background()); err != nil {
        log.Printf("Shutdown request error: %v", err)
    }
    close(server.doneChan)
  }()
}
