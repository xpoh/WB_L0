FROM golang:1.18-alpine as builder

WORKDIR /go/src/WB_L0

COPY . .

RUN go get -d -v ./...

RUN go build -o /app/orders ./cmd/.

FROM alpine:latest

COPY --from=builder /app/orders /app/orders

COPY website /app/website/

WORKDIR /app

ENTRYPOINT ["/app/orders"]