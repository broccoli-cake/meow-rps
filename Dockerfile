FROM golang:1.24.5

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o server

EXPOSE 8000

CMD ["./server"]

