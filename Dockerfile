FROM golang:1.24-alpine AS builder

RUN apk update && \
    apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o skillsRockTest ./cmd/main.go

FROM alpine:3.17

RUN apk update && \
    apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /app/skillsRockTest .
COPY --from=builder /app/migrations ./migrations

COPY --from=builder /app/.env .

EXPOSE 3000

CMD ["./skillsRockTest"]
