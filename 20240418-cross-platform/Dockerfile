FROM docker.io/library/golang:1.22 AS builder

COPY . .

RUN CGO_ENABLED=0 go build -o /build/hello hello.go

FROM scratch

EXPOSE 8080
COPY --from=builder /build/hello /usr/local/bin/hello

CMD ["/usr/local/bin/hello"]
