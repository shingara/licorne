FROM golang:1.4

RUN mkdir -p /go/bin ;\
  go get github.com/tools/godep


EXPOSE 9000
RUN mkdir -p /go/src/licorne
WORKDIR /go/src/licorne

ADD Godeps /go/src/licorne/Godeps
RUN godep restore
RUN go install -v github.com/onsi/ginkgo/ginkgo

ADD . /go/src/licorne
