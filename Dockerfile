FROM golang:1.20 AS builder
WORKDIR /build
COPY . ./
RUN go get -d -v && \
    CGO_ENABLED=0 go build -v -a -installsuffix cgo -o ./my-ip && \
    chmod +x ./my-ip

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build ./
EXPOSE 8080
CMD ["./my-ip"]
