package hashserver

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


func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func newHashRequest(password string, t *testing.T) *http.Request {
    data := url.Values{}
    data.Set("password", password)
    request, err := http.NewRequest("POST", "/hash", strings.NewReader(data.Encode()))
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    checkError(err, t)

    return request
}

func TestHashHandlerValidRequest(t *testing.T) {
    server := NewServer(0)
    hasher := hashencode.Sha512HasherImp
    expectedHash := hasher.Hash([]byte("angryMonkey"))
    password := "angryMonkey"

    request:= newHashRequest(password, t)
    requestRecorder := httptest.NewRecorder()

    http.HandlerFunc(server.hashHandler).ServeHTTP(requestRecorder, request)

    if status := requestRecorder.Code; status != http.StatusOK {
      t.Errorf("Status code incorrect. Expected %d Got %d", http.StatusOK, status)
    }

    if got := requestRecorder.Body.String(); got != expectedHash {
      t.Errorf("Response body incorrect. Expected %s Got %s", expectedHash, got)
    }
}

func TestHashHandlerOnlyAcceptsPost(t *testing.T) {
    server := NewServer(0)
    request, err := http.NewRequest("GET", "/hash", nil)
    checkError(err, t)

    requestRecorder := httptest.NewRecorder()
    http.HandlerFunc(server.hashHandler).ServeHTTP(requestRecorder, request)

    if status := requestRecorder.Code; status != http.StatusMethodNotAllowed {
      t.Errorf("Status code incorrect. Expected %d Got %d", http.StatusMethodNotAllowed, status)
    }

    expected := `{"Message": "Method Not Allowed"}`
    if got := requestRecorder.Body.String(); got != expected {
      t.Errorf("Response body incorrect. Expected %s Got %s", expected, got)
    }
}

func TestStatsHandlerCount(t *testing.T) {
    server := NewServer(0)
    password := "angryMonkey"
    hashRequest:= newHashRequest(password, t)
    requestCount := 3

    hashRequestRecorder := httptest.NewRecorder()
    for i := 0; i < requestCount; i++ {
      http.HandlerFunc(server.hashHandler).ServeHTTP(hashRequestRecorder, hashRequest)
    }

    request, err := http.NewRequest("GET", "/stats", nil)
    checkError(err, t)
    requestRecorder := httptest.NewRecorder()
    http.HandlerFunc(server.statsHandler).ServeHTTP(requestRecorder, request)

    type Body struct {
    		Total    int            `json:"total"`
    		Average  int            `json:"average"`
      }

    response := new(Body)
    body, _ := ioutil.ReadAll(requestRecorder.Body)
    json.Unmarshal(body, &response)

    if got := response.Total; got != requestCount {
      t.Errorf("Response body incorrect. Expected %d Got %d", requestCount, response.Total)
    }


}
