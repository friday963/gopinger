FROM golang:rc-alpine3.13

RUN apk add --no-cache git
RUN apk add build-base

WORKDIR /app/pinger

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build . 

CMD ["./ping_destination"]