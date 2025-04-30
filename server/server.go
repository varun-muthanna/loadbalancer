package server

import (
	"net"
	"sync"
	"time"
)

type Server struct {
	address           string
	domain 			  string	
	activeConnections int
	isHealthy         bool
	Mutex             sync.Mutex
}

func NewServer(domain string, address string) *Server {
	return &Server{
		domain: domain,
		address: address,
		activeConnections: 0,
		isHealthy:         true,
	}
}

func (s *Server) IncrementConnection() {
	s.Mutex.Lock()
	s.activeConnections += 1
	s.Mutex.Unlock()
}

func (s *Server) DecrementConnections() {
	s.Mutex.Lock()
	s.activeConnections -= 1
	s.Mutex.Unlock()
}

func (s *Server) SetHealth(health bool) {
	s.Mutex.Lock()
	s.isHealthy = health
	s.Mutex.Unlock()
}

func (s *Server) CheckHealth(timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", s.address, timeout) //hostname:port , 127.0.0.1:9001
	//conn to read and write to server

	if err != nil {
		return false
	}

	conn.Close()
	return true
}

func (s *Server) GetAddress() string{
	return s.address
}

func (s *Server) GetDomain() string{
	return s.domain
}

func (s *Server) GetActiveConnections() int{
	return s.activeConnections
}

func (s *Server) GetHealth() bool{
	return s.isHealthy
}