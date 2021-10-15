package lib
import (
	"encoding/json"
	"net/http"
)

const URL_PLACE = "url"

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

func GET_URL_FOR_DOWNLOAD(w http.ResponseWriter, r *http.Request) string {
	return (get_url_string_from_uri(URL_PLACE, w, r))[0]
}

