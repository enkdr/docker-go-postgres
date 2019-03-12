ARG GO_VERSION=1.11

FROM golang:${GO_VERSION} AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch

COPY --from=builder /app/* ./app/

EXPOSE 8080
ENTRYPOINT ["/app/docker-go-postgres"]