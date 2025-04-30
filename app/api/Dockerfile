FROM golang:1.24.2-alpine AS builder
WORKDIR /go/src/github.com/trihatmaja/Azan-Schedule
COPY ./* .
RUN pwd && \
    ls && \
    go build -o azan_api ./api/main.go && \
    ls

FROM alpine:latest
WORKDIR /app

COPY --from=builder /go/src/github.com/trihatmaja/Azan-Schedule/azan_api .
EXPOSE 1234
CMD ["/app/azan_api"]