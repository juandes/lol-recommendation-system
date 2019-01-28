BUILD        ?= $(CURDIR)/build/
IMAGE 	:= juandes/lol-champions-recommender

test:
	GOPATH=$(GOPATH) go test ./...

go-build:
	GOPATH=$(GOPATH) GOBIN=$(BUILD) go install  ./cmd/...
	
run-service-local:
	go run cmd/service/main.go --trainingset=static/winning_teams.csv

run-service-docker:
	docker pull ${IMAGE}
	docker run -t -p 8080:8080 ${IMAGE}

docker:
	docker build -t ${IMAGE} .

run-docker:
	docker run --expose 8080 -i -t ${IMAGE} bash