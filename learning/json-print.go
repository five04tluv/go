package main

import (
	"encoding/json"
	"net/http"
)

type Basic struct {
	Name   string
	Status []string
}

func main() {
	http.HandleFunc("/hello-world", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	bar := Basic{"Hello", []string{"Status"}}

	js, err := json.Marshal(bar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
