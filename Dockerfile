FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go mod download
RUN go build -o main .

FROM scratch

COPY --from=builder /build .

ENTRYPOINT ["./main"]

EXPOSE 3000