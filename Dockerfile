FROM golang:1.14.6-alpine

CMD ["ash"]
WORKDIR /app

# App
COPY go.mod go.sum /app/
RUN go mod download
COPY . /app/
RUN go build
