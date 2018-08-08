package server

import (
  "net/http"
  "net/http/httptest"
  "net/url"
  "encoding/json"
  "io/ioutil"
  "testing"
  "strings"
  "github.com/scottberke/password_hasher/hash"
)

// Function to fail test if any of the calls throw an error
func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

// Helper to build a new hash request
func newHashRequest(password string, t *testing.T) *http.Request {
    data := url.Values{}
    data.Set("password", password)

    request, err := http.NewRequest("POST", "/hash", strings.NewReader(data.Encode()))
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    checkError(err, t)

    return request
}

func TestHashHandlerValidRequest(t *testing.T) {
    // Create a new server with a port and a delay time of zero
    done := make(chan bool)
    server := NewServer(8080, 0, done)
    password := "angryMonkey"
    hasher := hashencode.Sha512HashEncoder
    expectedHash := hasher.Hash([]byte(password))

    request:= newHashRequest(password, t)
    requestRecorder := httptest.NewRecorder()

    http.HandlerFunc(server.hashHandler).ServeHTTP(requestRecorder, request)

    // Make sure our status code is correctly returned
    if status := requestRecorder.Code; status != http.StatusOK {
      t.Errorf("Status code incorrect. Expected %d Got %d", http.StatusOK, status)
    }
    // Make sure our hashed password is correctly returned
    if got := requestRecorder.Body.String(); got != expectedHash {
      t.Errorf("Response body incorrect. Expected %s Got %s", expectedHash, got)
    }
}

func TestHashHandlerOnlyAcceptsPost(t *testing.T) {
    // Create a new server with a port and a delay time of zero
    done := make(chan bool)
    server := NewServer(8080, 0, done)
    request, err := http.NewRequest("GET", "/hash", nil)
    checkError(err, t)

    requestRecorder := httptest.NewRecorder()
    http.HandlerFunc(server.hashHandler).ServeHTTP(requestRecorder, request)

    // Make sure our status code is correctly returned
    if status := requestRecorder.Code; status != http.StatusMethodNotAllowed {
      t.Errorf("Status code incorrect. Expected %d Got %d", http.StatusMethodNotAllowed, status)
    }

    // Make sure our error message is correctly returned
    expected := `{"Message": "Method Not Allowed"}`
    if got := requestRecorder.Body.String(); got != expected {
      t.Errorf("Response body incorrect. Expected %s Got %s", expected, got)
    }
}

func TestStatsHandlerCount(t *testing.T) {
    done := make(chan bool)
    server := NewServer(8080, 0, done)
    password := "angryMonkey"
    hashRequest:= newHashRequest(password, t)
    requestCount := 3

    hashRequestRecorder := httptest.NewRecorder()
    // Make requestCount number of requests
    for i := 0; i < requestCount; i++ {
      http.HandlerFunc(server.hashHandler).ServeHTTP(hashRequestRecorder, hashRequest)
    }

    request, err := http.NewRequest("GET", "/stats", nil)
    checkError(err, t)

    requestRecorder := httptest.NewRecorder()
    http.HandlerFunc(server.statsHandler).ServeHTTP(requestRecorder, request)

    // Struct to unmarshal json response into
    type Body struct {
    		Total    int  `json:"total"`
    		Average  int  `json:"average"`
      }

    response := new(Body)
    body, _ := ioutil.ReadAll(requestRecorder.Body)
    json.Unmarshal(body, &response)

    if got := response.Total; got != requestCount {
      t.Errorf("Response body incorrect. Expected %d Got %d", requestCount, response.Total)
    }
}


func TestStatsHandlerCountNoRequests(t *testing.T) {
    done := make(chan bool)
    server := NewServer(8080, 0, done)

    request, err := http.NewRequest("GET", "/stats", nil)
    checkError(err, t)

    requestRecorder := httptest.NewRecorder()
    http.HandlerFunc(server.statsHandler).ServeHTTP(requestRecorder, request)

    // Struct to unmarshal json response into
    type Body struct {
    		Total    int  `json:"total"`
    		Average  int  `json:"average"`
      }

    response := new(Body)
    body, _ := ioutil.ReadAll(requestRecorder.Body)
    json.Unmarshal(body, &response)

    if got := response.Total; got != 0 {
      t.Errorf("Response body incorrect. Expected %d Got %d", server.requestCount, response.Total)
    }
}
