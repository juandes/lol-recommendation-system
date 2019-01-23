
FROM golang:1.11

WORKDIR /go/src/github.com/juandes/lol-recommendation-system
COPY . .

RUN echo $GOPATH
RUN make test
RUN make go-build

FROM debian:jessie
COPY --from=0 /go/src/github.com/juandes/lol-recommendation-system/build/service /lol-recommendation-service/
COPY --from=0 /go/src/github.com/juandes/lol-recommendation-system/static/champions_key.json /data/
COPY --from=0 /go/src/github.com/juandes/lol-recommendation-system/static/winning_teams.csv /data/
COPY --from=0 /go/src/github.com/juandes/lol-recommendation-system/run.sh /app/

WORKDIR /app

CMD [ "/app/run.sh"]
