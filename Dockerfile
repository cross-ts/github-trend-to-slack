FROM golang:alpine AS builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/build
RUN apk add git && \
    go get -u github.com/golang/dep/cmd/dep
COPY Gopkg.* ./
RUN dep ensure --vendor-only
COPY *.go ./
RUN go build -a --installsuffix cgo -ldflags="-s -w" -o run *.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/build/run /run
CMD ["/run"]
