FROM golang:1.18 AS build_base
WORKDIR /go/src/finance

# Force the go compiler to use modules
ENV GO111MODULE=on

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# This image builds the weavaite server
FROM build_base AS builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM alpine
EXPOSE 8888
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/finance/main /app/
COPY --from=builder /go/src/finance/.env /app/
RUN apk add --no-cache curl
WORKDIR /app
CMD ["./main"]
