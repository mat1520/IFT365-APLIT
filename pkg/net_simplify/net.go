package netsimplify

import (
	"fmt"
	"net/http"
)

func CreateServer(addr string) *http.Server {
	return &http.Server{Addr: addr}
}

func StartServer(srv *http.Server) error {
	if srv == nil {
		return fmt.Errorf("error: el servidor es nil")
	}
	return srv.ListenAndServe()
}

func StopServer(srv *http.Server) error {
	if srv == nil {
		return fmt.Errorf("error: el servidor es nil")
	}
	return srv.Close()
}
