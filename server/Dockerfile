FROM golang:1.17.2

WORKDIR /go/src

COPY . .

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64

RUN go get github.com/pilu/fresh
CMD ["fresh"]