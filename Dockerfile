FROM golang:alpine

ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk update && apk add bash && apk add git && mkdir concur_download 
# python 
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python
RUN python3 -m ensurepip
RUN pip3 install --no-cache --upgrade pip setuptools

COPY . /concur_download/

RUN ls

WORKDIR /concur_download

RUN ls

RUN go get github.com/gorilla/mux && go get github.com/gorilla/handlers
RUN python3 -m venv env && source env/bin/activate
RUN pip3 install -r requirements.txt

EXPOSE 8080

CMD go run main.go
