package main

import (
	"encoding/json"
	"net/http"
)

type Foo struct {
	Hello  string
	Status string
}

func main() {
	http.HandleFunc("/hello-world", bar)
	http.ListenAndServe(":80", nil)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := Foo{Hello: "World!", Status: "OK"}
	json.NewEncoder(w).Encode(u)
}
