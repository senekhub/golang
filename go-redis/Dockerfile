FROM golang:alpine
RUN mkdir /app
ADD main.go /app/
WORKDIR /app
RUN go build -o main .
RUN adduser -S -D -H -h /app appuser
USER appuser
CMD ["./main"]