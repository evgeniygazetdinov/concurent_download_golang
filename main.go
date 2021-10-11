package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"os/exec"
	"bytes"
	server "lib"
)

func exec_command(args string, w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("date")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	fmt.Println("out")
	cmd.Stderr = &errb
	fmt.Println("err")
	err := cmd.Run()
	fmt.Println("run")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("call handler")
	server.SUCCESS_HANDLER(fmt.Sprint("1"), w, r )
	fmt.Println("after handler")

}

func uploader(w http.ResponseWriter, r *http.Request) {
	var fileUrl = "myfile";
	fmt.Println("Download 7 : " + fileUrl)
		
	log.Println("before for")
	for i := 1; i < 2; i++{
		fmt.Println("Downloaded %d:i ", i)
        go exec_command(fileUrl, w, r);
    }
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/null/", uploader)
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Println("listening on 8080")
	http.ListenAndServe(":8080", loggedRouter)

}
