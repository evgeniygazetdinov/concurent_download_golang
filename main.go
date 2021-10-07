package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"fmt"
	"io"
	"os"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	
)

const one_predicate = "a"
const two_predicate = "b"

func succes_handler(result string, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"Value": result, "Status": "true", "Err": ""})
}

func get_variable_string_from_uri(variable_name string, w http.ResponseWriter, r *http.Request) []string {
	variable, ok := r.URL.Query()[variable_name]
	if !ok || len(variable[0]) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "missed"}`))
	}
	return variable
}

func get_integers_for_equalize(one_name string, second_name string, w http.ResponseWriter, r *http.Request) (int, int) {

	one := get_variable_string_from_uri(one_name, w, r)
	two := get_variable_string_from_uri(second_name, w, r)
	first_int, err := strconv.Atoi(one[0])
	second_int, err := strconv.Atoi(two[0])
	if err != nil {
		fmt.Println((err))
	}
	return first_int, second_int
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func mul(w http.ResponseWriter, r *http.Request) {
	fileUrl := "https://golangcode.com/logo.svg"

	fmt.Println("Download 7 : " + fileUrl)
	for i := 1; i < 7; i++{
		fmt.Println("Downloaded %d:i ", i)
		fmt.Println(fileUrl)
        go DownloadFile("logo.svg", fileUrl)
    }
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/null/", mul)
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Println("listening on 8080")
	http.ListenAndServe(":8080", loggedRouter)
}
