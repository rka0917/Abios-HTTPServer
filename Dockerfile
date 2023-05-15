FROM golang:alpine

WORKDIR /Abios-HTTPServer

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /AbiosHTTPServer

EXPOSE 8080

CMD ["/AbiosHTTPServer"]
