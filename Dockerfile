FROM golang:1.14.9-alpine3.12 AS builder

WORKDIR src/github.com/tormath1/inframap-service

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
RUN go build -o inframap-web main.go

FROM alpine:3.12
COPY --from=builder /go/src/github.com/tormath1/inframap-service/inframap-web .
CMD ["/inframap-web"]
