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

func get_variable_string_from_uri(variaurble_name string, w http.ResponseWriter, r *http.Request) []string {
	// TODO ADD RE FOR SEARCH YOUTUBE URL
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

