package proxy

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/varun-muthanna/loadbalancer/balancer"
	"github.com/varun-muthanna/loadbalancer/forwardproxy"
)

type handler struct {
	addr string
	lb   *balancer.Balancer
	fp   *forwardproxy.ForwardProxy
}

func StartProxy(listenaddress string, lb *balancer.Balancer, fp *forwardproxy.ForwardProxy) {

	h := &handler{
		addr: listenaddress,
		lb:   lb,
		fp:   fp,
	}

	s := &http.Server{
		Addr:    listenaddress,
		Handler: h,
	}

	err := s.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to start proxy listener: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	defer s.Shutdown(ctx)

}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv := h.lb.GetLeastConnections(h.fp)
	//fmt.Printf("\n Active connections: %v",srv.GetActiveConnections())

	if srv == nil {
		log.Printf("No healthy backend servers available for connection from %s", h.addr)
		return
	}

	srv.IncrementConnection()
	defer srv.DecrementConnections()

	url := "http://" + srv.GetAddress() + r.RequestURI

	req, err := http.NewRequest(r.Method, url, r.Body)

	if err != nil {
		log.Printf("Error in creating request %s", err)
		return
	}

	req.Header = r.Header.Clone()

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf(" error making http request: %s", err)
		srv.SetHealth(false)
		return
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body) //captures w.Write([]) from backend servers

	if err != nil {
		log.Printf("error reading response: %s", err)
		return
	}

	w.Write(resBody) //essentially doing it twice 
	fmt.Println(string(resBody))
	
	//time.Sleep(2*time.Second)

}

// architecture

// client ----------- proxy ---------  server
//      make request        select
// 							connect
//		transfer data		transfer data
