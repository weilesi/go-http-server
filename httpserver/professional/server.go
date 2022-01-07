package professional

/**
  @Description: 封装了http-server,目前实现的是常规路由，后续添加static router和parameter router功能
  @author jun.hai
  @date 2021年12月29日 下午16:40:18
*/

import (
	"context"
	"net/http"
	"time"
)

type IServer interface {
	Routable
	Start(address string) error

	Shutdown(ctx context.Context) error
}

type structHttpServer struct {
	Name    string
	handler Handler
	filter  Filter
}

func (s *structHttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := NewContext(writer, request)
	s.filter(ctx)
}

func (s *structHttpServer) Route(method string, pattern string, handlerFunc handlerFunc) {
	s.handler.Route(method, pattern, handlerFunc)
}

func (s *structHttpServer) Start(address string) error {
	return http.ListenAndServe(address, s)
}

func (s *structHttpServer) Shutdown(ctx context.Context) error {
	time.Sleep(time.Second)
	return nil
}

// NewHttpServer 创建新的Server，给调用方暴露该函数
func NewHttpServer(name string, filters ...FilterBuilder) IServer {
	handler := NewTreeRouteHandler()
	var root Filter = handler.ServeHTTP

	for i := len(filters) - 1; i >= 0; i-- {
		b := filters[i]
		root = b(root)
	}

	server := &structHttpServer{
		Name:    name,
		handler: handler,
		filter:  root,
	}

	return server
}
