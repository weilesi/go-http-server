package sdk

/**
  @Description: 封装了http-server，陆续封装@TODO
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
	root    Filter
}

func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		ctx := NewContext(writer, request)
		s.root(ctx)
	})

	return http.ListenAndServe(address, nil)
}

func NewSdkHttpServer(name string, builders ...FilterBuilder) {

}
