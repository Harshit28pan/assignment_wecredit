FROM golang:1.17-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o main .
CMD ["./main"]
