export tag=v1.0
root:
	export ROOT=github.com/weilesi/go-http-server

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t weilesi/httpserver:${tag} .

push: release
	echo "pushing weilesi/httpserver"
	docker push weilesi/httpserver:${tag}
