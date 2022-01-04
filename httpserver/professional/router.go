package professional

/**
  @Description: 封装了路由树
  @author jun.hai
  @date 2021年12月28日 下午09:49:08
*/

import (
	"net/http"
	"strings"
)

type Routable interface {
	Route(method string, pattern string, handlerFunc handlerFunc)
}

type TreeRouteHandler struct {
	root *node
}

type node struct {
	path     string
	children []*node
	handler  handlerFunc
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0, 2),
	}
}

// NewTreeRouteHandler 需要实现Handler的方法
func NewTreeRouteHandler() Handler {
	return &TreeRouteHandler{
		root: &node{},
	}
}

// 从路由树中找对应节点，找到就执行，找不到提示Not Found
func (t *TreeRouteHandler) ServeHTTP(ctx *Context) {
	handler, found := t.findRouter(ctx.R.URL.Path)
	if !found {
		ctx.W.WriteHeader(http.StatusNotFound)
		_, _ = ctx.W.Write([]byte("Not Found"))
		return
	}
	handler(ctx)
}

// Route 创建新的路由树
func (t *TreeRouteHandler) Route(method string, pattern string, handlerFunc handlerFunc) {
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")

	cur := t.root
	for index, path := range paths {
		matchedChild, found := t.findMatchChild(cur, path)
		if found {
			cur = matchedChild
		} else {
			t.createNewNodeTree(cur, paths[index:], handlerFunc)
			return
		}
	}

	cur.handler = handlerFunc
}

//查找当前path对应的路由地址
func (t *TreeRouteHandler) findRouter(path string) (handlerFunc, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := t.root
	for _, p := range paths {
		matchChild, found := t.findMatchChild(cur, p)
		if !found {
			return nil, false
		}
		cur = matchChild
	}

	if cur.handler == nil {
		return nil, false
	}
	return cur.handler, true
}

func (t *TreeRouteHandler) findMatchChild(cur *node, p string) (*node, bool) {
	for _, child := range cur.children {
		if child.path == p {
			return child, true
		}
	}
	return nil, false
}

//在路由树上添加子节点集，例如 /info/address
func (t *TreeRouteHandler) createNewNodeTree(root *node, paths []string, handlerFunc handlerFunc) {
	cur := root
	for _, path := range paths {
		newNode := newNode(path)
		cur.children = append(cur.children, newNode)
		cur = newNode
	}
	cur.handler = handlerFunc
}
