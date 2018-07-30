package hashserver

import (
  "fmt"
  "log"
  "net/http"
  "time"
  "github.com/scottberke/password_hasher/hash"
)

func hashHandler(w http.ResponseWriter, r *http.Request) {
  password := r.FormValue("password")
  hasher := hashencode.Sha512HasherImp
  base64encodedSha512hash := hasher.Hash([]byte(password))

  time.Sleep(time.Duration(0)*time.Second) // TODO update this to 5 seconds per requirments
  fmt.Fprintf(w, "Hello World! Password: %s, Hash: %s", password, base64encodedSha512hash)
}



func HashServer() error {
    http.HandleFunc("/hash", hashHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))

    return nil
}
