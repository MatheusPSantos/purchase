FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o app .

FROM golang:1.23

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8888

CMD ["./app"]