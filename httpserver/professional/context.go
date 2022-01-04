package professional

/**
  @Description: 封装了http请求中的上下文
  @author jun.hai
  @date 2021年12月28日 下午2:02:06
*/
import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

// NewContext 创建上下文
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W: w,
		R: r,
	}
}

func (ctx *Context) ReadJson(data interface{}) error {
	body, err := io.ReadAll(ctx.R.Body)
	if nil != err {
		return err
	}
	return json.Unmarshal(body, data)
}

func (ctx *Context) WriteJson(status int, data interface{}) error {
	bt, err := json.Marshal(data)
	if nil != err {
		return err
	}
	_, err = ctx.W.Write(bt)
	if nil != err {
		return err
	}

	ctx.W.WriteHeader(status)
	return nil
}

// OkJson 封装常用的http状态码:200的返回
func (ctx *Context) OkJson(data interface{}) error {
	return ctx.WriteJson(http.StatusOK, data)
}

// InternalServerErrorJson 封装常用的http状态码:500的返回
func (ctx *Context) InternalServerErrorJson(data interface{}) error {
	return ctx.WriteJson(http.StatusInternalServerError, data)
}

// BadRequestJson 封装常用的http状态码:400的返回
func (ctx *Context) BadRequestJson(data interface{}) error {
	return ctx.WriteJson(http.StatusBadRequest, data)
}
