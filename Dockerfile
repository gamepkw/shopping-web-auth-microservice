FROM golang:1.21 AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

FROM alpine:latest

RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder /app/main .

COPY ./config/config.json .

EXPOSE 80

ENV TZ=Asia/Bangkok

CMD ["./main"]