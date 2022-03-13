package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-http-server/httpserver/metrics"
	"go-http-server/httpserver/professional"
	"go-http-server/httpserver/simple"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	simpleServerType       = "simple"
	professionalServerType = "professional"
	metricsServerType      = "metrics"
)

func main() {
	//serverType:1、启动简单服务传simple；2、启动封装后工程化的服务传professional;3、启动metrics服务
	callServerStart("metrics")
}

//HTTP-Server的简单实现
func simpleServer() {
	serverStartLog("simple-server", "simple")

	server := simple.NewSimpleHttpServer("simple")
	server.Route("/", simple.RootHandler)
	server.Route("/healthz", simple.HealthHandler)
	err := server.Start(":8088", false)
	if err != nil {
		log.Printf("Starting Simple Server:", err.Error())
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

//HTTP-Server的监控实现
func metricsServer() {
	metrics.Register()
	mux := http.NewServeMux()

	mux.HandleFunc("/", simple.RootHandler)
	mux.HandleFunc("/images", images)
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/healthz", simple.HealthHandler)
	if err := http.ListenAndServe(":8099", mux); err != nil {
		log.Fatalf("start http server failed, error: %s\n", err.Error())
	}
}

func images(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	w.Write([]byte(fmt.Sprintf("<h1>%d<h1>", randInt)))
}

func serverStartLog(serverName string, serverType string) {
	flag.Set("V", "4")
	log.Printf("Starting the " + serverType + " server." + " Server Name is " + serverName)
}

func callServerStart(serverType string) {
	switch serverType {
	case simpleServerType:
		simpleServer()
	case professionalServerType:
		professionalServer()
	case metricsServerType:
		metricsServer()
	}
}
