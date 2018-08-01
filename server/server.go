package hashserver

import (
  "fmt"
  "log"
  "net/http"
  "time"
  "os"
  "os/signal"
  "syscall"
  "context"
  "encoding/json"
  "github.com/scottberke/password_hasher/hash"
)

func (server *hashServer) hashHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    w.WriteHeader(405)
    w.Write([]byte(`{"Message": "Method Not Allowed"}`))
  } else {
    start := time.Now()
    password := r.FormValue("password")
    hasher := hashencode.Sha512HasherImp
    base64encodedSha512hash := hasher.Hash([]byte(password))
    elapsed := time.Since(start)
    server.TotalTime += elapsed
    server.RequestCount += 1
    log.Printf("method: %s", r.Method)
    log.Printf("hash_server: hashing password")
    time.Sleep(server.ResDelaySeconds*time.Second) // TODO update this to 5 seconds per requirments
    // fmt.Fprintf(w, "%v -- %v Requests: %i -- Hello World! Password: %s, Hash: %s", int64(server.TotalTime / time.Microsecond),  (int64(server.TotalTime) / int64(server.RequestCount)) , server.RequestCount, password, base64encodedSha512hash)
    fmt.Fprintf(w, "%s", base64encodedSha512hash)
  }
}

func (server *hashServer) shutdownHandler(w http.ResponseWriter, r *http.Request) {
  log.Printf("Shutting down server.")
  go func() {
    server.shutdown <- true
  }()
}

func (server *hashServer) statsHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet {
    w.WriteHeader(405)
    w.Write([]byte(`{"Message": "Method Not Allowed"}`))
  } else {
    res := make(map[string]int)
    res["total"] = server.RequestCount
    res["average"] = int(int64(server.TotalTime / time.Microsecond) / int64(server.RequestCount))

    data, _ := json.Marshal(res)
    w.Header().Set("Content-Type", "application/json")
    w.Write(data)
  }
}

func (server *hashServer) ShutdownServer() {
  stop := make(chan os.Signal, 1)
  signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

  select {
  case sig := <-stop:
     log.Printf("Shutdown request (signal: %v)", sig)
     // time.Sleep(6*time.Second)
  case sig := <-server.shutdown:
     log.Printf("Shutdown request (/shutdown %v)", sig)
     // time.Sleep(6*time.Second)
  }

  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()

  if err := server.Shutdown(ctx); err != nil {
      log.Printf("Shutdown request error: %v", err)
  }

}

type hashServer struct {
  http.Server
  shutdown      chan bool
  RequestCount  int
  TotalTime     time.Duration
  ResDelaySeconds  time.Duration
}

func NewServer(resDelaySeconds time.Duration) *hashServer {
    mux := http.NewServeMux()
    server := &hashServer{
      Server:   http.Server{Addr: ":8080",
                            Handler: mux,
                          ReadTimeout: 10 * time.Second,
                          WriteTimeout: 10 * time.Second,},
      shutdown: make(chan bool),
      RequestCount: 0,
      ResDelaySeconds: resDelaySeconds,
    }

    mux.HandleFunc("/hash", server.hashHandler)
    mux.HandleFunc("/shutdown", server.shutdownHandler)
    mux.HandleFunc("/stats", server.statsHandler)

    return server
}
