FROM golang:1.20 AS builder
WORKDIR /build
COPY . ./
RUN go get -d -v && \
    go build -v -o ./my-ip && \
    chmod +x ./my-ip

FROM debian:11
COPY --from=builder /build /app
RUN ls -la /app
WORKDIR /app
EXPOSE 8080
CMD ["./my-ip"]
