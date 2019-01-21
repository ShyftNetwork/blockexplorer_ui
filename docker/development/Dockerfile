FROM golang:1.10.3-alpine

RUN apk add --update git make gcc musl-dev linux-headers ca-certificates

COPY ./ /go/src/github.com/ShyftNetwork/blockexplorer_ui/shyft_api
WORKDIR /go/src/github.com/ShyftNetwork/blockexplorer_ui/shyft_api

CMD go run -v *.go

EXPOSE 8080
