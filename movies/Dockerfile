FROM golang

RUN mkdir -p /go/src/github.com/ranggakusuma/go-kit-example

ADD . /go/src/github.com/ranggakusuma/go-kit-example

RUN go get  -t -v ./...
RUN go build -o movies cmd/main.go

EXPOSE 8080

CMD ["./movies"]
