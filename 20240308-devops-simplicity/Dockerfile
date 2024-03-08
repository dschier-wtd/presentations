FROM docker.io/library/golang:1.22 AS build

ENV CGO_ENABLED=0
COPY . .
RUN go build -v -o /build/hello cmd/hello.go

FROM scratch

COPY --from=build /build/hello /usr/local/bin/hello
EXPOSE 3000/tcp
CMD ["/usr/local/bin/hello"]
