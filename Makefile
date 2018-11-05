BUILD        ?= $(CURDIR)/build/
IMAGE 	:= juandes/lol-champions-recommender

build:
	go build
	./lol_example

run:
	go run cmd/lol_example/main.go

run_simple:
	go run cmd/simple_example/main.go

run_assessment:
	go run cmd/system_assessment/main.go

test:
	GOPATH=$(GOPATH) go test ./...

go-build:
	GOPATH=$(GOPATH) GOBIN=$(BUILD) go install  ./cmd/...

docker:
	docker build -t ${IMAGE} .

run-docker:
	docker run --expose 8080 -i -t ${IMAGE} bash

run-service:
	docker run -t -p 8080:8080 juandes/lol-champions-recommender