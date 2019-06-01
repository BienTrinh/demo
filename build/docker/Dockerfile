FROM golang:latest
LABEL maintainer="bien@trustingsocial.com"

RUN go get -u github.com/golang/dep/cmd/dep
# use golangci-lint v1.15.0 for golang 1.10.x compatible
RUN go get -u github.com/golangci/golangci-lint/cmd/golangci-lint \
    && cd $GOPATH/src/github.com/golangci/golangci-lint/cmd/golangci-lint \
    && git checkout v1.15.0 \
    && go build \
    && cp -f golangci-lint $GOPATH/bin


ENV APP_HOME  /go/src/github.com/BienTrinh/demo

WORKDIR $APP_HOME
COPY Gopkg.lock Gopkg.toml $APP_HOME/
RUN dep ensure --vendor-only -v

COPY . $APP_HOME/