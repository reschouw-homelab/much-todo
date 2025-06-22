FROM golang:1.24.4 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o much-todo-backend .

FROM alpine:latest

COPY --from=builder /app/much-todo-backend /app/much-todo-backend

ENTRYPOINT ["/app/much-todo-backend"]

EXPOSE 8080