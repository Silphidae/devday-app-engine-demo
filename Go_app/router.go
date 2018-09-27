package main

import (
	"net/http"
)

func router(w http.ResponseWriter, r *http.Request) {
	var head = r.URL.Path
	switch head {
	case "/api":
		routeAPI(w, r)
	default:
		http.Error(w, "", http.StatusNotFound)
		return
	}
}

func routeAPI(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		routeCandiesList(w, r)
		return
	case http.MethodPost:
		routeCandyCreate(w, r)
		return
	default:
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
}
