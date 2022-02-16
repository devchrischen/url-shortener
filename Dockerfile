FROM golang:1.17

ENV GO111MODULE="on"
# ENV DB_HOST="host.docker.internal"
# ENV REDIS_HOST="host.docker.internal"

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o ./app

CMD ["./app"]

EXPOSE 8080