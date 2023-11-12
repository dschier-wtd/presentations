FROM docker.io/library/golang:1.20 AS build

WORKDIR /usr/local/src
ENV CGO_ENABLED=0

COPY src/ /usr/local/src/

RUN go mod tidy \
    && go get . \
    && go build -o /usr/local/bin/hello -v .

FROM scratch

COPY --from=build /usr/local/bin/hello /usr/local/bin/

EXPOSE 3000/tcp

CMD ["/usr/local/bin/hello"]
