FROM golang:1.4

RUN go get github.com/robfig/revel; go get github.com/robfig/revel/revel; go get gopkg.in/mgo.v2

EXPOSE 9000
RUN mkdir -p /go/src/licorne
WORKDIR /go/src/licorne
ADD . /go/src/licorne
