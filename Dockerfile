FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY *.yaml ./

RUN go build -o /ruuvigw-go

CMD [ "/ruuvigw-go" ]