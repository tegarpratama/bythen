FROM golang:1.22-alpine

WORKDIR /go/src/app

COPY /app/go.mod /app/go.sum ./
RUN go mod download

COPY ./app .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]