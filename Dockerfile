FROM golang

COPY . /go/src/github.com/OGFris/SluxDB

RUN go get -d -t /go/src/github.com/OGFris/SluxDB/cmd/sluxdb

RUN go install github.com/OGFris/SluxDB/cmd/sluxdb

ENTRYPOINT /go/bin/sluxdb

EXPOSE 6060
