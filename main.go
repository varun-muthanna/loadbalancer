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
	"github.com/varun-muthanna/loadbalancer/forwardproxy"
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
	var bannedDomains = []string{"*1.com",}

	for _, m := range cfg.BackendServers {
		for key , value := range m{
			servers = append(servers, server.NewServer(key,value))
		}
	}

	lb := balancer.NewLoadBalancer(servers)
	fp := forwardproxy.NewForwardProxy(bannedDomains)

	healthInterval := time.Duration(cfg.HealthCheckInterval) * time.Second
	healthTimeout := time.Duration(cfg.HealthCheckTimeout) * time.Second

	health.StartHealthCheck(servers, healthInterval, healthTimeout)

	fmt.Printf("Load Balancer listening on: %s", cfg.ListenAddress)
	proxy.StartProxy(cfg.ListenAddress, lb,fp)

}
