FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test -v -vet=off ./tests/...
RUN cd cmd && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../analytics .

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/analytics .

CMD ["./analytics"]