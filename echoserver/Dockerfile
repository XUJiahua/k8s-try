# https://docs.docker.com/develop/develop-images/multistage-build/
FROM golang:1.14.6
WORKDIR /echoserver
COPY main.go .
COPY go.mod .
RUN CGO_ENABLED=0 GOOS=linux go build -o echoserver .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /echoserver/echoserver .
CMD ["./echoserver"]  
