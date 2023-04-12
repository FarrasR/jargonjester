FROM golang:1.20.1-alpine

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
RUN mkdir /app
WORKDIR /app

COPY . .
COPY .env .env
RUN go mod download

RUN go build -o /discord cmd/discord/main.go

EXPOSE 8888
CMD [ "/discord" ]
