FROM golang:1.20 AS builder
WORKDIR /build
COPY . ./
RUN go get -d -v && \
    go build -v -o ./my-ip && \
    chmod +x ./my-ip

# FIXME
#FROM alpine:latest
#WORKDIR /app
#COPY --from=builder /build ./
EXPOSE 8080
CMD ["./my-ip"]
