package health

import (
	"log"
	"time"

	"github.com/varun-muthanna/loadbalancer/balancer"

	"github.com/varun-muthanna/loadbalancer/server"
)

func StartHealthCheck(lb *balancer.Balancer, interval, timeout time.Duration) { //both of type time.Duration
	//create new ticker , (timeout of tcp connection) for each server

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	go func() { // separately runnning to main program , other sequential execution 
		for range ticker.C {
			for _, srv := range lb.Servers {
				go func(s *server.Server) { // go routine in for loop , for closure as it may always use last value of loop always (pass by reference) , now we pass by value
					isHealthy := s.CheckHealth(timeout)
					s.SetHealth(isHealthy)
					if !isHealthy {
						log.Printf("Server failed, Address :%s", s.Address)
					}
				}(srv)
			}
		}
	}()
}
