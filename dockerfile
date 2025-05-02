FROM golang:alpine AS builder

LABEL authors="websalad" \
    com.datadoghq.ad.logs='[{"source": "go", "service": "categories"}]'

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o main cmd/main.go

CMD ["./main"]