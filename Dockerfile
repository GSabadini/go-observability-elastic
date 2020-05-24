FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

#RUN apk update
#RUN apk upgrade
#RUN apk add --update go=1.8.3-r0 gcc=6.3.0-r4 g++=6.3.0-r4

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o main

FROM scratch
COPY --from=builder /build .
ENTRYPOINT ["./main"]
EXPOSE 3000