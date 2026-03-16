FROM golang:1.21

WORKDIR /app

COPY . .
RUN go mod tidy

COPY . .

RUN go build -o main .

EXPOSE 8081

CMD ["./main"]
