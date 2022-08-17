FROM golang:alpine3.16 as build
ADD . /src
WORKDIR /src
RUN go build
FROM alpine:3.16
COPY --from=build /src/qulog2loki /usr/bin
CMD ["/usr/bin/qulog2loki"]
