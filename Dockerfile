FROM golang:1.10.4
WORKDIR /go/src/github.com/ilgooz/service-http-server
COPY . .
RUN go install ./...
CMD httpserver --serverAddr :2300