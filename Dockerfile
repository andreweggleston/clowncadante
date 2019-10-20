FROM golang:1.11

USER nobody

RUN mkdir -p /go/src/github.com/openshift/golang-ex
COPY . /go/src/github.com/andreweggleston/clowncadante
WORKDIR /go/src/github.com/andreweggleston/clowncadante
RUN CGO_ENABLED=0 go build

CMD ["./clowncadante"]
