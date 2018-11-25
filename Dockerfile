FROM golang:1.11.2 as builder

WORKDIR /go/src/brr/
COPY . .

RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go install -a ./...


FROM alpine:3.6 as runner

WORKDIR /root/
COPY ./db ./db
COPY --from=builder /go/bin /usr/bin
