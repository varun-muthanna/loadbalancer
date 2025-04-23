package proxy

import (
	"io"
	"log"
	"net"

	"github.com/varun-muthanna/loadbalancer/balancer"
)

func StartProxy(listenaddress string, lb *balancer.Balancer) {

	//listen for incoming client connections
	listener, err := net.Listen("tcp", listenaddress)

	if err != nil {
		log.Fatalf("Failed to start proxy listener: %v", err)
	}

	defer listener.Close()

	for {

		//accept incoming requests
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("Failed to accept client connection: %v", err)
			continue
		}

		//new go routine to handle client connections asynchonously
		go HandleConnecton(conn, lb)
	}
}

func HandleConnecton(clientConn net.Conn, lb *balancer.Balancer) {

	defer clientConn.Close()

	srv := lb.GetLeastConnections()

	if srv == nil {
		log.Printf("No healthy backend servers available for connection from %s", clientConn.RemoteAddr())
		return
	}

	//increment and dec , create conn with server ,  transfer data

	srv.IncrementConnection()
	defer srv.DecrementConnections()

	srvconn, err := net.Dial("tcp", srv.Address)

	if err != nil {
		log.Printf("Failed to connect to backend server %s: %v", srv.Address, err)
		srv.SetHealth(false)
		return
	}

	defer srvconn.Close()

	go func() {
		_ , err :=io.Copy(srvconn,clientConn)
		if err !=nil{
			log.Printf("Error proxying data from client to server: %v", err)
		}
	}() // full duplex communication :- simulatenously handle server-client communication  

	_, err = io.Copy(clientConn, srvconn)

	if err != nil {
		log.Printf("Error proxying data from server to client: %v", err)
	}

}

// architecture

// client ----------- proxy ---------  server
//      make request        select
// 							connect
//		transfer data		transfer data
