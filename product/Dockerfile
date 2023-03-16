FROM golang:1.19.5-bullseye as builder
RUN mkdir -p /build/go-mesh-consul/product
COPY ./go.* /build/
COPY ./go-mesh-consul/product /build/go-mesh-consul/product

WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app /build/go-mesh-consul/product/main.go 

FROM alpine:3.17.1
WORKDIR /app
COPY --from=builder /build/ /app
EXPOSE 9090
ENTRYPOINT ["./app"]