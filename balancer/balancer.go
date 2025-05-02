package balancer

import (
	"sync"

	"github.com/varun-muthanna/loadbalancer/server"
	"github.com/varun-muthanna/loadbalancer/forwardproxy"
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

func (lb *Balancer) GetLeastConnections(fp *forwardproxy.ForwardProxy) *server.Server {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	var mx = 1 << 15
	var res *server.Server

	for _, srv := range lb.servers { //index ,server
		srv.Mutex.Lock()

		if srv.GetHealth() && srv.GetActiveConnections() < mx {

			var c bool = fp.IsBanned(srv)

			if !c{
				res = srv
				mx = srv.GetActiveConnections()
			}

		}

		srv.Mutex.Unlock()
	}

	// str := fmt.Sprintln( "Domain:" ,res.GetDomain(),"Active Connections" ,res.GetActiveConnections())
	// fmt.Println(str)

	return res
}

