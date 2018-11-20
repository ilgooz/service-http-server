FROM golang:1.10.4
WORKDIR /go/src/github.com/ilgooz/service-website
COPY . .
RUN go install ./...
CMD website