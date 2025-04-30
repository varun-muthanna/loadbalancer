# Load Balancer

A simple load balancer project written in Go.

## Features
- Least Connection request distribution
- Health checks

## Usage
go run main.go

go run backend.go <port number>  // get the servers running , make sure same as config.json

Send continuous tcp requests to lb port to check balancer work 
