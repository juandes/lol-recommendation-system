BUILD        ?= $(CURDIR)/build/
IMAGE 	:= juandes/lol-champions-recommender

run-service-local:
	go run cmd/service/main.go --trainingset=static/winning_teams.csv

test:
	GOPATH=$(GOPATH) go test ./...

go-build:
	GOPATH=$(GOPATH) GOBIN=$(BUILD) go install  ./cmd/...

docker:
	docker build -t ${IMAGE} .

run-docker:
	docker run --expose 8080 -i -t ${IMAGE} bash

run-service:
	docker pull ${IMAGE}
	docker run -t -p 8080:8080 ${IMAGE}