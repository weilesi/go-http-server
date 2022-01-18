package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"go-http-server/httpserver/professional"
	"go-http-server/httpserver/simple"
)

const (
	simpleServerType       = "simple"
	professionalServerType = "professional"
)

func main() {
	//serverType:1、启动简单服务传simple；2、启动封装后工程化的服务传professional
	callServerStart("professional")
}

//HTTP-Server的简单实现
func simpleServer() {
	serverStartLog("simple-server", "simple")

	server := simple.NewSimpleHttpServer("simple")
	server.Route("/", simple.RootHandler)
	server.Route("/healthz", simple.HealthHandler)
	err := server.Start(":8088", false)
	if err != nil {
		glog.Errorf("Starting Simple Server:", err.Error())
	}
}

//HTTP-Server工程化封装实现
func professionalServer() {
	serverStartLog("user", "professional")

	server := professional.NewHttpServer("pro-server", professional.MetricFilterBuilder)
	server.Route("GET", "/", professional.UserIndexHandler)
	server.Route("POST", "/user/login", professional.UserLoginHandler)

	err := server.Start(":8099")
	if err != nil {
		fmt.Println("Service startup encountered an error：" + err.Error())
		panic(err)
	}
}

func serverStartLog(serverName string, serverType string) {
	flag.Set("V", "4")
	glog.Info("Starting the " + serverType + " server." + " Server Name is " + serverName)
}

func callServerStart(serverType string) {
	switch serverType {
	case simpleServerType:
		simpleServer()
	case professionalServerType:
		professionalServer()
	}
}
