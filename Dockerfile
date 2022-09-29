FROM golang:latest AS builder

RUN mkdir /build

ADD . /build

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags '-extldflags "-static"' -o main cmd/main.go

FROM alpine

COPY --from=builder /build/main /app/

EXPOSE 8080

WORKDIR /app

CMD ["./main"]
