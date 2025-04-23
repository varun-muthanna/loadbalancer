package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync/atomic"
)

var connectionCount int32

func handleConnection(conn net.Conn) {
	atomic.AddInt32(&connectionCount, 1)
	defer func() {
		atomic.AddInt32(&connectionCount, -1)
		conn.Close()
	}()

	addr := conn.RemoteAddr().String()
	log.Printf("Connection established from %s", addr)

	// Echo server: send back any received data
	_, err := io.Copy(conn, conn)
	if err != nil {
		log.Printf("Connection error with %s: %v", addr, err)
	}

	log.Printf("Connection closed from %s", addr)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run backend.go <port>")
		return
	}
	port := os.Args[1]
	address := ":" + port

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}
	defer listener.Close()

	log.Printf("Backend server listening on %s", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}
