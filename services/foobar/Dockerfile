FROM golang:1.11 as builder
WORKDIR /module
COPY ./go.mod /module/go.mod
COPY ./go.sum /module/go.sum
RUN go mod download
COPY . /module
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app

FROM alpine:3.10
WORKDIR /root/
COPY --from=builder /module/bin .
EXPOSE 3000

ENV GIN_MODE release
CMD ./app
