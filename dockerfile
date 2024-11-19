FROM golang:1.21-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .


WORKDIR /app/cmd
RUN go build -o main .


EXPOSE 8181

CMD ["./main"]