FROM golang:1.20

WORKDIR /app

COPY . ./

RUN go get -d -v
RUN go build -v -o my-ip

COPY /app/my-ip /app/my-ip

EXPOSE 8080

CMD ["/app/my-ip"]
