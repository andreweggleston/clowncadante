FROM golang:1.8 AS build
COPY . /go/src/github.com/andreweggleston/clowncadante
WORKDIR /go/src/github.com/andreweggleston/clowncadante
RUN CGO_ENABLED=0 go build

FROM scratch
COPY --from=build /go/src/github.com/andreweggleston/clowncadante/clowncadante /clowncadante
ENTRYPOINT ["/clowncadante"]
