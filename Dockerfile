FROM golang:1.20.7-alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go get ./...
RUN go build -o /go/bin/server ./src

FROM alpine
COPY --from=builder /go/bin/server /app/server
# COPY --from=builder /go/src/config.yaml /app

WORKDIR /app
EXPOSE 5050
CMD ["./server"]