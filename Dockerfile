FROM golang:1.23.4-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest

CMD [ "air", "-c", ".air.toml" ]