FROM golang:1.4

RUN mkdir -p /go/bin ;\
  go get github.com/tools/godep



# /* RUN go get github.com/codegangsta/negroni; \ */
# /*  go get github.com/gorilla/context; \ */
# /*  go get github.com/gorilla/mux; \ */
# /*  go get gopkg.in/mgo.v2 */
#
# /* RUN go get github.com/onsi/ginkgo/ginkgo ;\ */
# /*  go get github.com/onsi/gomega */

EXPOSE 9000
RUN mkdir -p /go/src/licorne
WORKDIR /go/src/licorne

ADD Godeps /go/src/licorne/Godeps
RUN godep restore
RUN go install -v github.com/onsi/ginkgo/ginkgo

ADD . /go/src/licorne
