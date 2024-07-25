FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .


FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/main /app/main

CMD ["/app/main"]
