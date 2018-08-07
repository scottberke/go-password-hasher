[![Build Status](https://travis-ci.org/scottberke/go-password-hasher.svg?branch=master)](https://travis-ci.org/scottberke/go-password-hasher)
# Password Hasher In Go

## Description
This application, when launched, will listen for HTTP requests on a provided port.
Three endpoints currently exist:
 1. [`/hash`](#hash)
 2. [`/stats`](#stats)
 3. [`/shutdown`](#shutdown)

These endpoints are documented below with example usage.

## Assumptions
1. My assumption is that the folks reviewing this repo know that my experience with Go and statically typed languages is limited and, despite a strong effort, this repo undoubtably contains idiomatic errors.
2. The requirements under section 2 `Hash and Encode Passwords over HTTP` do not state that the response should be JSON. The requirements under section 4 ` Statistics Endpoint` do explicitly state the response should be JSON. I took this to mean that the response for the `/hash` endpoint should return exactly what's shown - a text response containing the base64 encoded string of the SHA512 hash of the provided password
3. The requirements under section 4 `Statistics Endpoint` don't mention anything about persisting data. I took this to mean that I only needed to return the statistics for the current running server's uptime and the passwords that have been hashed during this time.

## To Run Locally
To install and build, execute:
```bash
  $ go get github.com/scottberke/go-password-hasher
  $ cd $GOPATH/src/github.com/scottberke/go-password-hasher
  $ go build
```

To see available flags:
```bash
$ ./go-password-hasher -h
Usage of ./go-password-hasher:
  -delay int
    	number of seconds to delay hash response (default 5)
  -port int
    	a port to start the server on (default 8080)
```

To run:
```bash
$ cd $GOPATH/src/github.com/scottberke/go-password-hasher
$ ./go-password-hasher -port=8080 -delay=0
```

To run tests:
```bash
$ cd $GOPATH/src/github.com/scottberke/go-password-hasher
$ go test ./...
```

## Endpoints
### Hash
#### POST /hash
Use to retrieve a SHA512 hash that has been base64 encoded of a plain text password.

##### Request
```bash
curl -X POST \
  http://localhost:8080/hash \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -d password=angryMonkey
```
##### Response 200 OK (Content-Type: text/plain; charset=utf-8)
```
ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==
```

### Stats
#### GET /stats
Use to retrieve stats for the server since it's been running. Stats include the total number of passwords that have been hashed and the average time it has taken to hash and encode each password.

##### Request
```bash
curl -X GET \
  http://localhost:8080/stats
```
##### Response 200 OK (Content-Type: application/json)
```json
{"average":43,"total":4}
```

### Shutdown
#### GET /shutdown
Use to gracefully shutdown the server.

##### Request
```bash
curl -X GET \
  http://localhost:8080/shutdown
```
##### Response 200 OK
```json
  {"Message": "Shutdown in progress"}
```
