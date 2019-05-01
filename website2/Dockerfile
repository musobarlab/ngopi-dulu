FROM golang:1.11.1-alpine3.8

WORKDIR /usr/app

# Add the source code
ENV SRC_DIR=/usr/app/

ADD . $SRC_DIR

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -o app .

EXPOSE 8001

ENTRYPOINT ["sh", "-c", "./app"]