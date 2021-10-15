package lib
import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
)

func SUCCESS_HANDLER(result string, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"Value": result, "Status": "true", "Err": ""})
}

func get_url_string_from_uri(variable_name string, w http.ResponseWriter, r *http.Request) []string {
	// TODO ADD RE FOR SEARCH YOUTUBE URL
	variable, ok := r.URL.Query()[variable_name]
	if !ok || len(variable[0]) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "url missed"}`))
	}
	return variable
}

func get_url_for_download(url_string string, w http.ResponseWriter, r *http.Request) string {
	raw_url := get_url_string_from_uri(url_string, w, r)
	url_link, err := strconv.Atoi(raw_url)
	if err != nil {

		fmt.Println((err))
	}
	return url_link
}

