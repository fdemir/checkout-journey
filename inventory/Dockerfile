FROM golang:1.21.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY *.go ./

RUN CGO_ENABLED=1 GOOS=linux go build -o /inventory

EXPOSE 3001

CMD [ "/inventory" ]