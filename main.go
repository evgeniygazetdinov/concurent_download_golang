package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const one_predicate = "a"
const two_predicate = "b"


// func DownloadFile() error {

// 	// Get the data
// 	ytdl := youtube_dl.YoutubeDl{}
// 	ytdl.Path = "$GOPATH/src/app/srts" // for example
// 	err := ytdl.DownloadVideo(noeTdQaBYCk)
// 	if err != nil {
// 		log.Printf("%v", err)
// 	}
// }

func mul(w http.ResponseWriter, r *http.Request) {
	fileUrl := "https://golangcode.com/logo.svg"

	fmt.Println("Download 7 : " + fileUrl)
	for i := 1; i < 2; i++{
		fmt.Println("Downloaded %d:i ", i)
        // go DownloadFile()
    }
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/null/", mul)
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Println("listening on 8080")
	http.ListenAndServe(":8080", loggedRouter)

}
