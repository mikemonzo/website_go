FROM golang:1.23.4-alpine

WORKDIR /app

COPY . /app/

RUN go install github.com/air-verse/air

RUN go mod download

CMD [ "air", "-c", ".air.toml" ]