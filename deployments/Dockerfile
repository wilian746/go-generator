FROM golang:1.14.2 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o go-generator cmd/go-generator/main.go

FROM alpine:3
COPY --from=builder /build/go-generator /bin/
COPY .bashrc /root/.bashrc
RUN apk update && apk add --no-cache git bash openssh curl
CMD ["./bash"]
