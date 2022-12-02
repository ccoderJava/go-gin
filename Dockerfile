FROM golang:latest
MAINTAINER ccoderJava "congccoder@gmail.com"

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /go-gin

COPY . /go-gin

RUN go build .

EXPOSE 8000

ENTRYPOINT ["./go-gin"]