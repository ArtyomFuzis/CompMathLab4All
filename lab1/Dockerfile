FROM golang:1.23.6-alpine3.21
WORKDIR /app
COPY . .
RUN go build -o backend witchmat
CMD ["./backend"]