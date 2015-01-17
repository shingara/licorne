FROM golang:1.4

RUN go get github.com/robfig/revel; go get github.com/robfig/revel/revel

EXPOSE 9000
RUN mkdir -p /go/src/licorne
WORKDIR /go/src/licorne
ADD . /go/src/licorne
