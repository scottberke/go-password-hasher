# Password Hasher In Go

## Description
This application sets up a server with three endpoints - `/hash`, `/stats` and `/shutdown`.

The `/hash` endpoint accepts POST requests and consumes a plain text password. The endpoint returns, in plain text, the sha512 hash of the password that has been base64 encoded.

The `/stats` endpoint accepts GET requests and returns JSON containing the total number of requests that have been made to the server and the average time it's taken to hash each password request since the server has been started.

The `/shutdown` endpoint accepts GET requests and will gracefully shutdown the server.

These endpoints are illustrated below with example usage.

## Assumptions
1. The requirements for this assignment indicate that Go was to be used as the language - It's my assumption that the folks reviewing this repo know that my experience with Go and statically typed languages is limited and, despite a strong effort, this repo undoubtably contains idiomatic errors.
2. The requirements under section 2 `Hash and Encode Passwords over HTTP` do not state that the response should be JSON. The requirements under section 4 ` Statistics Endpoint` do explicitly state the response should be JSON. I took this to mean that the response for the `/hash` endpoint should return exactly what's shown - a text response containing the base64 encoded string of the SHA512 hash of the provided password
3. The requirements under section 4 `Statistics Endpoint` don't mention anything about persisting data. I took this to mean that I only needed to return the statistics for the current running server's uptime and the passwords that have been hashed during this time.
4. The requirements under section 4 `Statistics Endpoint` are vague as to how the average time it has taken to process all of the requests should be calculated. Since the `/hash` endpoint is supposed to have the socket open for 5 seconds, the average time would always be around 5 seconds. My assumption was that this stat was reporting on the average time to hash and encode the passwords in the course of the request so the average time being returned is the average time in microseconds the application takes to hash and encode t   


## To Run Locally

## Endpoints
### Hash
#### POST /hash
Use to retrieve a sha512 hash that has been base64 encoded of a plain text password.

##### Request
```bash
curl -X POST \
  http://localhost:8080/hash \
  -H 'Cache-Control: no-cache' \
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
  http://localhost:8080/stats \
  -H 'Cache-Control: no-cache' \
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
