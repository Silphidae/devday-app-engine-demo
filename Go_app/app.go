package main

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", router)
	appengine.Main()
}

func mustJSON(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
