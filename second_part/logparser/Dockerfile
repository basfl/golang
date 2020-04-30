FROM golang:1.10 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main .


FROM alpine:latest AS production
COPY  --from=builder /app .
RUN mkdir /data 
#RUN touch /data/temp.txt
VOLUME "/data/"
CMD ["./main"]