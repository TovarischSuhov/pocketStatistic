FROM golang:latest AS builder

WORKDIR /usr/src/app

RUN ls -l

RUN go build -mod=vendor -o getPocketStatistic cmd/main.go

COPY --from=build getPocketStatistic getPocketStatistic

EXPOSE 8080

ENTRYPOINT ["getPocketStatistic"]
