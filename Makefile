SHELL := /bin/sh

all: schema dev

test:
	go test -v -race `go list ./... | grep -v internal/docker | grep -v proto`

cover:
	go test --race `go list ./... | grep -v /vendor | grep -v /cmd/erply ` -coverprofile cover.out.tmp && \
		cat cover.out.tmp | grep -v "bindata.go" > cover.out && \
		go tool cover -func cover.out && \
		rm cover.out.tmp && \
		rm cover.out

schema:
	cd internal/storage/postgres/schema && go generate

dev:
	docker-compose -f docker/docker-compose.yaml up
