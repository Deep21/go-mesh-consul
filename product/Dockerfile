FROM golang:1.19.5-bullseye as builder
RUN mkdir /build
RUN go mod init github.com/deep21
ADD . /build/
WORKDIR /build
ENTRYPOINT ["tail"]
CMD ["-f"]

# RUN CGO_ENABLED=0 GOOS=linux go build -a -o customer .

# FROM alpine:3.17.1
# WORKDIR /app
# COPY --from=builder /build/ /app
# EXPOSE 9091
# ENTRYPOINT ["tail"]
# CMD ["-f"]
