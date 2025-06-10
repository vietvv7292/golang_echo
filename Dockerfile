FROM golang:1.21

WORKDIR /app

COPY ./app/go.mod ./
RUN go mod download

COPY ./app .

RUN go build -o server .

CMD ["./server"]
