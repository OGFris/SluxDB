FROM golang

COPY . /go/src/github.com/OGFris/SluxDB

RUN go get -d -v -t /go/src/github.com/OGFris/SluxDB/cmd/sluxdb && go install github.com/OGFris/SluxDB/cmd/sluxdb

ENTRYPOINT /go/bin/sluxdb

EXPOSE 6060
