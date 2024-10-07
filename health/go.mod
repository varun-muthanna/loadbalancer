module github.com/varun-muthanna/loadbalancer/health

replace github.com/varun-muthanna/loadbalancer/balancer v0.0.0 => ../balancer

replace github.com/varun-muthanna/loadbalancer/server v0.0.0 => ../server

go 1.22.1

require github.com/varun-muthanna/loadbalancer/balancer v0.0.0

require github.com/varun-muthanna/loadbalancer/server v0.0.0
