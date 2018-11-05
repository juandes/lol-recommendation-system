
FROM golang:1.11

WORKDIR /go/src/github.com/juandes/knn-recommender-system
COPY . .

RUN echo $GOPATH
RUN make test
RUN make go-build

FROM debian:jessie
COPY --from=0 /go/src/github.com/juandes/knn-recommender-system/build/service /lol-recommender-service/
COPY --from=0 /go/src/github.com/juandes/knn-recommender-system/data/champions_key.json /data/
COPY --from=0 /go/src/github.com/juandes/knn-recommender-system/data/winning_teams.csv /data/
COPY --from=0 /go/src/github.com/juandes/knn-recommender-system/run.sh /app/

WORKDIR /app

CMD [ "/app/run.sh"]
