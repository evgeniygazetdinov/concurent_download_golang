package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"os/exec"
	server "lib"
)


func exec_command(url string, nums chan string,w http.ResponseWriter, r *http.Request){
	out, err := exec.Command("bash", "-c", fmt.Sprintf("youtube-dl %s", url)).Output()
	if err != nil {
        panic(url)
    }
	nums <-string(out);
}

func uploader(w http.ResponseWriter, r *http.Request) {
	var url = fmt.Sprintf("https://www.youtube.com/watch?v=%s",server.GET_URL_FOR_DOWNLOAD(w, r))
	my_channel := make(chan string) 
    go exec_command(url, my_channel, w, r);
	server.SUCCESS_HANDLER(<-my_channel, w, r);
	
	close(my_channel)

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/downloader/", uploader);

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Println("listening on 8080")
	http.ListenAndServe(":8081", loggedRouter)

}
