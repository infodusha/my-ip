FROM golang:1.20

WORKDIR /app

COPY . ./

RUN go get -d -v
RUN go build -v -o my-ip

EXPOSE 8080

CMD ["/my-ip"]
