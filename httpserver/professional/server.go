package professional

/**
  @Description: 封装了http-server
  @author jun.hai
  @date 2021年12月29日 下午16:40:18
*/

import "net/http"

type Routable interface {
	Route(method string, pattern string, handlerFunc handlerFunc)
}

type Server interface {
	Routable
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler Handler
	filter  Filter
}

func (s *sdkHttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := NewContext(writer, request)
	s.filter(ctx)
}

func (s *sdkHttpServer) Route(method string, pattern string, handlerFunc handlerFunc) {
	s.handler.Route(method, pattern, handlerFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, s)
}

func NewSdkHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewTreeRouteHandler()
	var root Filter = handler.ServeHTTP

	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}

	server := &sdkHttpServer{
		Name:    name,
		handler: handler,
		filter:  root,
	}

	return server
}
