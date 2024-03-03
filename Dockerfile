FROM golang:1.22.0-alpine3.19 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o server 

FROM alpine:3.19.1

WORKDIR /app

COPY --from=builder /app/server /app

EXPOSE 8080 

CMD ["./server"] 
