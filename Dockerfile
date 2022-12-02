FROM scratch
MAINTAINER ccoderJava "congccoder@gmail.com"

WORKDIR /go-gin

COPY . /go-gin

EXPOSE 8000

CMD ["./go-gin"]