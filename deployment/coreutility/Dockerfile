# Deps stage
FROM golang:alpine3.16 AS deps
ADD ./coreutility /app/coreutility
ADD ./common /app/common
WORKDIR /app/coreutility
RUN go mod download

# Build stage
FROM golang:alpine3.16 AS builder
RUN apk update && apk add curl && apk add coreutils && apk add gcc libc-dev
ADD ./coreutility /app/coreutility
ADD ./common /app/common
COPY --from=deps /go /go
WORKDIR /app/coreutility
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags '-linkmode=external' -o /build/main /app/coreutility/cmd/main.go

# Final stage
FROM golang:alpine3.16
RUN apk --no-cache add curl ca-certificates tzdata
COPY --from=builder /app/coreutility/config /app/config
COPY --from=builder /build /app/main
RUN chmod +x /app/main
ENTRYPOINT ["/app/main"]
EXPOSE 8013