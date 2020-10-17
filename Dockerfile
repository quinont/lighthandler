from golang:1.15.3-buster as builder
WORKDIR /go/src/app
COPY main.go .
RUN go get
RUN CGO_ENABLED=0 go build -o /tmp/app

from alpine:3.12.0
WORKDIR /myapp
COPY --from=builder /tmp/app .
CMD ["./app"]
