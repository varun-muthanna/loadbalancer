package balancer

import (
	"sync"

	"github.com/varun-muthanna/loadbalancer/server"
)

type Balancer struct {
	Servers []*server.Server //slice for pointers
	Mutex   sync.Mutex
}

func NewLoadBalancer(srv []*server.Server) *Balancer{
	return &Balancer{
		Servers: srv,
	}
}

func (lb *Balancer) GetLeastConnections() *server.Server {
	lb.Mutex.Lock()
	defer lb.Mutex.Unlock()

	var mx = 1 << 15
	var res *server.Server

	for _, srv := range lb.Servers { //index ,server
		srv.Mutex.Lock()

		if srv.IsHealthy && srv.ActiveConnections < mx {
			res = srv
			mx = srv.ActiveConnections
		}

		srv.Mutex.Unlock()

	}

	return res

}
