FROM golang:latest AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN ls -l

RUN go build -mod=vendor -o getPocketStatistic cmd/main.go

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/getPocketStatistic /app/getPocketStatistic

EXPOSE 8080

ENTRYPOINT ["getPocketStatistic"]
