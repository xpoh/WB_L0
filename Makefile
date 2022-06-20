GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
BINARY_NAME=url
VERSION?=0.0.1
SERVICE_PORT?=8000
DOCKER_REGISTRY?= #if set it should finished by /
EXPORT_RESULT?=false # for CI please set EXPORT_RESULT to true

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test build vendor

ddos:
	go run ./test/ddos.go

wb:
	go run ./cmd/wb_l0.go

stress:
	echo "GET http://localhost:8080/id/avNfxTCEXjHNCVQSMMFNSZEex" | vegeta attack -duration=5s | tee results.bin | vegeta report
