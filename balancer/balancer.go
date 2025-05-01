package balancer

import (
	"sync"

	"github.com/varun-muthanna/loadbalancer/server"
)

type Balancer struct {
	servers []*server.Server //slice for pointers
	mutex   sync.Mutex
	bannedDomains []string 
}

func NewLoadBalancer(srv []*server.Server) *Balancer{
	return &Balancer{
		servers: srv,
	}
}

func NewLoadBalancerWithForwardProxy(srv []*server.Server, domains[]string) *Balancer{
	return &Balancer{
		servers: srv,
		bannedDomains:domains,
	}
}

func (lb *Balancer) GetLeastConnections() *server.Server {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	var mx = 1 << 15
	var res *server.Server

	//fmt.Println(len(lb.servers))

	for _, srv := range lb.servers { //index ,server
		srv.Mutex.Lock()

		if srv.GetHealth() && srv.GetActiveConnections() < mx {
			c := true  //tied to outside loop

			for _, d := range lb.bannedDomains{
				if srv.GetDomain() == d {
					c = false 
				}
			}

			if c{ //if continue then lock is not free (blocked)
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

