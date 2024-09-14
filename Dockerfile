# FROM golang:1.21.0

# ENV GIN_MODE release

# WORKDIR /go/src/app

# RUN go install github.com/air-verse/air@latest

# COPY ./app .

# CMD air

FROM golang:1.21-alpine

WORKDIR /go/src/app

# RUN go install github.com/air-verse/air@latest

COPY ./app .

RUN go mod download

RUN go build -o /main .

EXPOSE 8080

CMD ["./main"]