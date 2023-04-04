FROM golang:1.20

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go get -d -v
RUN go build -v -o ./my-ip
RUN chmod +x ./my-ip

EXPOSE 8080

CMD ["/app/my-ip"]
