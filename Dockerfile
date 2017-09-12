FROM golang:1.6
RUN go get -t -v "github.com/stretchr/testify"
COPY . $GOPATH/src/github.com/Cardfree/teamcity-sdk-go
WORKDIR $GOPATH/src/github.com/Cardfree/teamcity-sdk-go
RUN go test -v -c ./types && \
    go test -v -c ./teamcity
CMD ./run-test.sh
