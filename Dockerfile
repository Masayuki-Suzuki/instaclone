FROM golang:1.14.4

WORKDIR /go/src/instaclone
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh && \
    go get github.com/go-sql-driver/mysql && \
    go get github.com/gorilla/mux && \
    go get github.com/gorilla/handlers && \
    go get firebase.google.com/go && \
    go get github.com/spf13/viper

CMD ["fresh", "-c", "fresh.conf.yml"]
