package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/varun-muthanna/loadbalancer/balancer"
	"github.com/varun-muthanna/loadbalancer/config"
	"github.com/varun-muthanna/loadbalancer/health"
	"github.com/varun-muthanna/loadbalancer/proxy"
	"github.com/varun-muthanna/loadbalancer/server"  //go walks the directory tree and resolves it normally.
)

func main() {
	//read the flag , default if no flag then config.json 
	configpath := flag.String("config", "config.json", "Path to configuration file") //checks current working
	flag.Parse()

	cfg, err := config.LoadConfig(*configpath)

	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	var servers []*server.Server

	for _, addr := range cfg.BackendServers {
		servers = append(servers, server.NewServer(addr))
	}

	lb := balancer.NewLoadBalancer(servers)

	healthInterval := time.Duration(cfg.HealthCheckInterval) * time.Second
	healthTimeout := time.Duration(cfg.HealthCheckTimeout) * time.Second

	health.StartHealthCheck(lb, healthInterval, healthTimeout)

	fmt.Printf("Load Balancer listening on: %s", cfg.ListenAddress)
	proxy.StartProxy(cfg.ListenAddress, lb)

}
