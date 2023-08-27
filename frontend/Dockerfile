FROM golang:1.21.0-alpine3.18
WORKDIR /Chat
COPY *.html ./
COPY *.css ./
COPY main.go .
COPY go.mod .
RUN go mod download
RUN go build -o main main.go
CMD ["./main"]



