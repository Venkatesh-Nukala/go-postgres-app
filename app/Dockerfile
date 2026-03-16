FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod init app
RUN go mod tidy

RUN go build -o main .

EXPOSE 8081

CMD ["./main"]
