FROM golang:alpine

ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk update && apk add bash && apk add git && mkdir concur_download


COPY . /concur_download/

RUN ls

WORKDIR /concur_download

RUN ls

RUN go get github.com/gorilla/mux && go get github.com/gorilla/handlers && go get github.com/youtube-videos/go-youtube-dl
EXPOSE 8080

CMD go run main.go
