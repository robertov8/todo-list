FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o todo-list ./cmd/api

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/todo-list .

EXPOSE 4000

CMD ["./todo-list"]