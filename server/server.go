package server

import (
	"net"
	"sync"
	"time"
)

type Server struct {
	Address           string
	ActiveConnections int
	IsHealthy         bool
	Mutex             sync.Mutex
}

func NewServer(address string) *Server {
	return &Server{
		Address:           address,
		ActiveConnections: 0,
		IsHealthy:         true,
	}
}

func (s *Server) IncrementConnection() {
	s.Mutex.Lock()
	s.ActiveConnections += 1
	s.Mutex.Unlock()
}

func (s *Server) DecrementConnections() {
	s.Mutex.Lock()
	s.ActiveConnections -= 1
	s.Mutex.Unlock()
}

func (s *Server) SetHealth(health bool) {
	s.Mutex.Lock()
	s.IsHealthy = health
	s.Mutex.Unlock()
}

func (s *Server) CheckHealth(timeout time.Duration) bool {
	conn, err := net.DialTimeout("tcp", s.Address, timeout) //hostname:port , 127.0.0.1:9001
	//conn to read and write to server

	if err != nil {
		return false
	}

	conn.Close()
	return true
}
