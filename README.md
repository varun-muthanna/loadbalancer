A simple load balancer that use least connection algorithm in Golang


STEPS TO RUN 

go run main.go -- start the load balancer 
go run backend.go (port) as mentioned in config files , in this case 9001, 9002, 9003 . Will set up servers.

telnet localhost 8080 (Manually send simulate client connections to load balancer).
Simulate at least >5 connections and check the least connection algorithm work. 
