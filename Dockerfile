FROM golang:latest AS builder

WORKDIR /app
COPY . .

ENV USER=appuser
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/app" \
    --shell "/sbin/nologin" \
    --no-create-home \
    "${USER}"

RUN go mod download
RUN CGO_ENABLED=0 go build -o /app/bin/main

ENTRYPOINT ["/app/bin/main"]