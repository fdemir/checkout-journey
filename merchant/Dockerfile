FROM golang:1.21.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY *.go ./

RUN GOOS=linux go build -o /merchant

EXPOSE 3004

CMD [ "/merchant" ]