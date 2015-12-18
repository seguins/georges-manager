FROM golang:1.5

RUN apt-get update && apt-get install -y curl && apt-get clean
RUN curl -L https://github.com/docker/compose/releases/download/1.5.2/docker-compose-`uname -s`-`uname -m` > /bin/docker-compose
RUN chmod +x /bin/docker-compose

RUN mkdir -p /tmp/docker/

COPY . /go/src/github.com/seguins/georges-manager
WORKDIR /go/src/github.com/seguins/georges-manager/georges

RUN go get -d -v

CMD ["go", "run", "georges.go"]
