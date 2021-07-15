FROM golang:alpine as builder

RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN go build
RUN go build -o listrepo cmd/listrepo/main.go

FROM alpine
COPY --from=builder /build/go-wrk /usr/local/bin/go-wrk
COPY --from=builder /build/listrepo /usr/local/bin/listrepo
