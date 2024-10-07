module github.com/varun-muthanna/loadbalancer

replace github.com/varun-muthanna/loadbalancer/balancer v0.0.0 => ./balancer

replace github.com/varun-muthanna/loadbalancer/health v0.0.0 => ./health

replace github.com/varun-muthanna/loadbalancer/proxy v0.0.0 => ./proxy

replace github.com/varun-muthanna/loadbalancer/server v0.0.0 => ./server

replace github.com/varun-muthanna/loadbalancer/config v0.0.0 => ./config

go 1.22.1

require (
	github.com/varun-muthanna/loadbalancer/balancer v0.0.0
	github.com/varun-muthanna/loadbalancer/config v0.0.0
	github.com/varun-muthanna/loadbalancer/health v0.0.0
	github.com/varun-muthanna/loadbalancer/proxy v0.0.0
	github.com/varun-muthanna/loadbalancer/server v0.0.0
)
