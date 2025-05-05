package balancer

import (
	"sync"

	"github.com/varun-muthanna/loadbalancer/server"
)

type Balancer struct {
	servers []*server.Server //slice for pointers
	mutex   sync.Mutex
}

func NewLoadBalancer(srv []*server.Server) *Balancer{
	return &Balancer{
		servers: srv,
	}
}

func (lb *Balancer) GetLeastConnections(host string) *server.Server {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	var mx = 1 << 15
	var res *server.Server

	for _, srv := range lb.servers { //index ,server
		srv.Mutex.Lock()

		if srv.GetHealth() && srv.GetActiveConnections() < mx && host == srv.GetDomain(){
			res=srv
			mx=srv.GetActiveConnections()
		}

		srv.Mutex.Unlock()
	}

	// str := fmt.Sprintln( "Domain:" ,res.GetDomain(),"Active Connections" ,res.GetActiveConnections())
	// fmt.Println(str)

	return res
}

