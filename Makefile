export GO111MODULE ?= on
export GOPROXY ?= https://goproxy.cn,direct

default: server

full:
	swag init && go run ./main.go start-api --config config.yaml --log_level debug

server:
	go run ./main.go start-api --config config.yaml --log_level debug

swagger:
	swag init

.PHONY: default run server
