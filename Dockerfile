FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go get github.com/githubnemo/CompileDaemon

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

CMD /wait && CompileDaemon --build="go build main.go" --command="./main"