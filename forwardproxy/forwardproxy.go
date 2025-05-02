package forwardproxy

import (
	"strings"

	"github.com/varun-muthanna/loadbalancer/server"
)

type ForwardProxy struct{
	bannedDomains []string 
}

func NewForwardProxy (bannedDomains[] string) *ForwardProxy{
	return &ForwardProxy{
		bannedDomains : bannedDomains,
	}
}

func (f *ForwardProxy) IsBanned(srv *server.Server) bool {

	for _, d := range f.bannedDomains{

		if strings.HasPrefix(d,"*") && strings.HasSuffix(srv.GetDomain(),d[1:]){
			return true //suffix matching only is d startswith "*"
		}

		if srv.GetDomain() == d { //exact match 
			return true 
		}
	}
	
	return false 
}