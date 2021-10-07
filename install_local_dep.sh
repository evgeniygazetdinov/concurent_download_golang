eval "GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64: #&& go mod init dowmload_concur && go get github.com/gorilla/mux && go get github.com/gorilla/handlers && go get github.com/youtube-videos/go-youtube-dl";