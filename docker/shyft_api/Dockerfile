FROM golang:1.10.3-alpine

RUN apk add --update git make gcc musl-dev linux-headers ca-certificates

COPY ./shyft_api /go/src/github.com/ShyftNetwork/blockexplorer_ui/shyft_api
WORKDIR /go/src/github.com/ShyftNetwork/blockexplorer_ui/shyft_api

RUN go get -u github.com/kardianos/govendor
RUN govendor sync

CMD go run -v *.go

EXPOSE 8080
