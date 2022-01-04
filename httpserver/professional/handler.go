package professional

/**
  @Description: 封装了http请求处理接口
  @author jun.hai
  @date 2021年12月28日 下午21:35:18
*/

type Handler interface {
	ServeHTTP(ctx *Context)
	Routable
}

type handlerFunc func(ctx *Context)
