FROM golang:1.14
RUN mkdir -p /go/src/desafio
WORKDIR /go/src/desafio
COPY src/desafio/* /go/src/desafio/
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go test -v
RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o desafio