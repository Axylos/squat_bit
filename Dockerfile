FROM golang:1.12.1-alpine

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build app

CMD ["app"]
