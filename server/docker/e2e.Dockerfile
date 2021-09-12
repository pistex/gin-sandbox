FROM golang:1.17-alpine3.14

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN apk update && apk upgrade
RUN apk add --no-cache gcc musl-dev

CMD ["go", "test",  "./...",  "-v"]