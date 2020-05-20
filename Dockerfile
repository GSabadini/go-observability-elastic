FROM golang:1.14-stretch

WORKDIR /go-apm-elastic

COPY . .

RUN go mod download
RUN go get github.com/cespare/reflex

COPY reflex.conf /

EXPOSE 3000

ENTRYPOINT ["reflex", "-c", "/reflex.conf"]
