FROM golang:1-alpine
RUN apk add -u make git
RUN mkdir /go/github.com
RUN go get github.com/c-garcia/nanogo 

