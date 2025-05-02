# Load Balancer

A simple http load balancer project written in Go.

## Features
- Least Connection request distribution
- Forward proxy policy option with lb configuration 
- Health checks

## Usage
go run main.go

go run backend.go <port number>  // get the servers running , make sure same as config.json

Send continuous http requests to lb port to check balancer work 
ex: hey -n 100 -c 5 http://localhost:8080 (100 total requests 5 concurrent connections)
