# Reverse Proxy
A simple reverse proxy written in Go.

## Features
- Least Connection load balancing to local servers
- Host header based forwarding of requests
- Health checks of local servers
- Public DNS resolution 

## Usage
go run main.go

go run backend.go <port number>  // get the servers running , make sure same as config.json

Send continuous http requests to lb port to reverse proxy work with load balancing to local servers 

ex: wrk -t4 -c5 -d2s -s host.lua http://localhost:8080
