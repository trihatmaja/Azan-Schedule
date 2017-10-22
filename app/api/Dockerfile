FROM golang:1.9-stretch AS builder
WORKDIR /go/src/github.com/trihatmaja/Azan-Schedule
COPY ../../* .
RUN CGO_ENABLED=0 make api

FROM alpine:latest
WORKDIR /app

COPY --from=builder /go/src/github.com/trihatmaja/Azan-Schedule/azan_api .
EXPOSE 1234
CMD ["/app/azan_api"]