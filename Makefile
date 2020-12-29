CURRENT=$(shell basename $(shell pwd))

LDFLAGS += -X "v/$(CURRENT)/util.CommitID=$(shell git rev-parse master)"
LDFLAGS += -X "v/$(CURRENT)/util.Date=$(shell date +%FT%T%z)"
LDFLAGS += -X "v/$(CURRENT)/util.Tag=$(shell git describe --always --tags)"
LDFLAGS += -X "v/$(CURRENT)/util.Branch=$(shell git rev-parse --abbrev-ref HEAD)"

all: server client cron scheduler http

client:
	@echo make client
	@go build -o $(CURRENT)_client cmd/client/main.go
.PHONY: client

server:
	@echo make server
	@go build -o $(CURRENT)_server cmd/server/main.go
.PHONY: server

http:
	@echo make http
	@go build -o $(CURRENT)_http cmd/http/main.go
.PHONY: http

cron:
	@echo make cron
	@go mod tidy
	@go build -o $(CURRENT)_cron cmd/cron/main.go
.PHONY: cron

scheduler:
	@echo make scheduler
	@go mod tidy
	@go build -o $(CURRENT)_scheduler cmd/scheduler/main.go
.PHONY: scheduler

