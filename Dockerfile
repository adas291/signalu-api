FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN --mount=type=cache,target="/root.cache/go-build" CGO_ENABLED=1 GOOS=linux go build -o /signalai-api

FROM golang:1.22
WORKDIR /app
COPY --from=builder /signalai-api /app
CMD ["./signalai-api"]