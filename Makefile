.PHONY: run-containers
run-containers:
	# Ensure tabs are used here
	docker run --rm -d -p 9001:80 --name server1 kennethreitz/httpbin
	docker run --rm -d -p 9002:80 --name server2 kennethreitz/httpbin
	docker run --rm -d -p 9003:80 --name server3 kennethreitz/httpbin

## stop: stops all demo services
.PHONY: stop
stop:
	# Ensure tabs are used here
	docker stop server1
	docker stop server2
	docker stop server3

## run: starts demo http services
.PHONY: run-proxy-server
run-proxy-server:
	# Ensure tabs are used here
	go run cmd/main.go

## run-all: starts demo services and the proxy server
.PHONY: run-all
run-all: run-containers run-proxy-server
