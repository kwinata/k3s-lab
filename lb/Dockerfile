FROM golang:1.26-alpine

COPY . .
RUN go build -o lb
EXPOSE 8080

CMD ["./lb"]