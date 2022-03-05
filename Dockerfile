FROM golang:1.17

ENV GO111MODULE="on"

# For dev
# ENV DB_HOST="host.docker.internal"
# ENV REDIS_HOST="host.docker.internal"

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o app

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.5.2

RUN go install github.com/cosmtrek/air@v1.27.8

CMD ["./app"]

EXPOSE 8080