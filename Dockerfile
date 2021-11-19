FROM docker.io/golang:alpine

RUN apk add --update --no-cache alpine-sdk bash ca-certificates \
      libressl \
      tar \
      git openssh openssl yajl-dev zlib-dev cyrus-sasl-dev openssl-dev build-base coreutils

RUN apk add --no-cache curl git openssl \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
    && apk del curl

WORKDIR $GOPATH/src/github.com/hugomd/go-todo

COPY . .

RUN go mod init
RUN go mod tidy
RUN go mod vendor

CMD ["go", "run", "main.go"]
