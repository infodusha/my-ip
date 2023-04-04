FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN go get -d -v && go build -v -o ./my-ip

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app ./
RUN chmod +x ./my-ip
EXPOSE 8080
ENTRYPOINT ["./my-ip"]
