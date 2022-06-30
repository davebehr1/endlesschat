FROM golang:1.18-bullseye as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest AS production
COPY --from=builder /app .
EXPOSE 5003
CMD ["./main","serve"]