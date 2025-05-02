package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

var connectionCount int32

type handler struct {
	add string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	atomic.AddInt32(&connectionCount, 1)

	defer func() {
		atomic.AddInt32(&connectionCount, -1)
	}()

	var str string = fmt.Sprintf("hello you got loadbalanced to me at %s", h.add)

	w.Write([]byte(str))
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run backend.go <port>")
		return
	}
	port := os.Args[1]
	address := ":" + port

	h := &handler{
		add: address,
	}

	s := &http.Server{
		Addr:    address,
		Handler: h,
	}

	err := s.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	defer s.Shutdown(ctx)

	log.Printf("Backend server listening on %s", address)
}
