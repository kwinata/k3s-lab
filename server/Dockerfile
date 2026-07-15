FROM golang:1.26-alpine

COPY . .
RUN go build -o server
EXPOSE 8080

CMD ["./server"]