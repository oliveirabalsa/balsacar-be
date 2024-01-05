
FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o balsacar-be ./cmd

EXPOSE 8081

CMD ["./balsacar-be"]
