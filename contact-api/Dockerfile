FROM golang as builder
WORKDIR /build/api
COPY go.mod ./
RUN go mod download
COPY . ./
RUN set -x && \
    go get -d -v . && \
    GOOS=linux go build -o api

FROM alpine
WORKDIR /root
COPY --from=builder /build/api/api .
EXPOSE 3000
CMD ["./api"]