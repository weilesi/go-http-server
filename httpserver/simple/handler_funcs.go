package simple

import (
	"fmt"
	"github.com/golang/glog"
	"go-http-server/httpserver/utils"
	"io"
	"net/http"
	"os"
	"strconv"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RootServer handler")

	header := r.Header
	//将request 中带的 header 写入 response header
	for key, val := range header {
		w.Header().Set(key, arrayToString(val))
	}
	w.Header().Set("server-name", "My name is simple server.")
	//把系统环境变量Version写入到response的header中
	w.Header().Set("version", os.Getenv("VERSION"))

	ip := utils.GetClientIP()
	glog.Infof("客户端请求IP为：" + ip + ", Http返回码：" + strconv.Itoa(http.StatusOK))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HealthServer handler")
	fmt.Fprintf(w, "Hello,"+r.URL.Path[1:]+"\n")
	io.WriteString(w, "当前状态："+strconv.Itoa(http.StatusOK))
}

//遍历数组中所有元素追加成string
func arrayToString(arr []string) string {
	var result string
	for _, val := range arr {
		result += val
	}
	return result
}
