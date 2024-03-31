# Choose whatever you want, version >= 1.16
FROM golang:1.21.0

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest && go install github.com/go-delve/delve/cmd/dlv@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]