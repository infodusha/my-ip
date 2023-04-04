FROM golang:1.20
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN go get -d -v && \
    go build -v -o ./my-ip

FROM alpine:latest
WORKDIR /app
COPY --from=0 /app ./
RUN chmod +x ./my-ip
EXPOSE 8080
ENTRYPOINT ["/app/my-ip"]
