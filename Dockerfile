FROM golang:1.20

WORKDIR $GOPATH/src/my-ip

COPY . .

RUN go get -d -v ./...
RUN go build -v ./...

EXPOSE 8080

CMD ["my-ip"]
