FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /app/qr ./cmd/qr
RUN CGO_ENABLED=0 go build -o /app/health ./cmd/health

FROM gcr.io/distroless/static-debian12

COPY --from=builder /app/qr /app/qr
COPY --from=builder /app/health /app/health

HEALTHCHECK --interval=5s --timeout=3s --start-period=10s CMD /app/health
ENTRYPOINT ["/app/qr"]
