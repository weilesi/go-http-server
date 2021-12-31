package main

import (
	"flag"
	"github.com/golang/glog"
	"go-http-server/httpserver/simple"
)

func main() {
	simpleServer()

}

//HTTP-Server的简单实现
func simpleServer() {
	flag.Set("V", "4")
	glog.Info("Starting the simple server...")

	server := simple.NewSimpleHttpServer("simple")
	server.Route("/", simple.RootHandler)
	server.Route("/healthz", simple.HealthHandler)
	err := server.Start("localhost:8088")
	if err != nil {
		glog.Errorf("Starting Simple Server:", err.Error())
	}
}

//HTTP-Server更多的封装实现
func professionalServer() {

}
