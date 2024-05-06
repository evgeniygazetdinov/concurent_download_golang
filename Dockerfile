FROM golang:alpine

ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

  

COPY . /simple_crud/

RUN ls

WORKDIR /simple_crud

RUN ls

RUN go get github.com/gorilla/mux && go get github.com/gorilla/handlers && go get -u github.com/go-sql-driver/mysql

EXPOSE 8080

CMD go run main.go
