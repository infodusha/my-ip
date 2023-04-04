FROM golang:1.20 AS builder
WORKDIR /build
COPY . ./
RUN go get -d -v && \
    go build -v -o ./my-ip && \
    chmod +x ./my-ip

#FROM alpine:latest
#WORKDIR /root
#COPY --from=builder /build ./
EXPOSE 8080
CMD ["./my-ip"]
