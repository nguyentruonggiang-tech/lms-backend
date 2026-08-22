FROM golang:1.26.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o lms-backend ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/lms-backend .

CMD ["./lms-backend"]
