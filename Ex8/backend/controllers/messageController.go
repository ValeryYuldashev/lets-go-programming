package controllers

import (
	"encoding/json"
	"net/http"
)

var Message = func(w http.ResponseWriter, r *http.Request) {
	mgs := map[string]interface{}{"message": "Hello, world!"}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(mgs)
}
