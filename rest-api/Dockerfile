FROM golang:1.10 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go get -u github.com/gorilla/mux 
RUN go build  .


FROM alpine:latest AS production

COPY  --from=builder /app .
RUN chmod 755 cmd.sh
ENTRYPOINT ["/cmd.sh"]