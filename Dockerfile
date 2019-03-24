FROM arm64v8/golang:1.12.1-alpine

WORKDIR /go/src/app

COPY . .

RUN apk update \
&& apk add --virtual build-dependencies \
build-base \
gcc \
git \
wget

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build app

CMD ["app"]
