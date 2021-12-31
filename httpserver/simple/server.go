package simple

import (
	"net/http"
	"time"
)

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
	Shutdown() error
}

type simpleHttpServer struct {
	Name string
}

func NewSimpleHttpServer(name string) Server {
	return &simpleHttpServer{
		Name: name,
	}
}

func (s *simpleHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handlerFunc)
}

func (s *simpleHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func (s *simpleHttpServer) Shutdown() error {
	//因为简单服务，模拟一下就可以
	time.Sleep(5 * time.Second)

	return nil
}
