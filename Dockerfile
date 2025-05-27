FROM golang:1.24-bookworm AS base

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o lumbaumbah-backend

EXPOSE 8080

CMD ["/build/lumbaumbah-backend"]