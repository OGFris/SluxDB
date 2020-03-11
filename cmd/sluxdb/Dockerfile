FROM golang

COPY ../../ /go/src/github.com/OGFris/SluxDB

RUN go install github.com/OGFris/SluxDB/cmd/sluxdb

ENTRYPOINT /go/bin/sluxdb

EXPOSE 6060
