FROM golang:1.6
RUN go get -t -v "github.com/stretchr/testify"
COPY . /go/src/github.com/umweltdk/teamcity
WORKDIR /go/src/github.com/umweltdk/teamcity
RUN go test -v -c ./types && \
    go test -v -c ./teamcity
CMD ./run-test.sh