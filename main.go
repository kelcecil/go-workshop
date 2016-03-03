package main

import (
	"net/http"
)

var GAMES map[int]interface{}

func main() {
	http.Handle("/NewGame", NewGameHandler{})

	http.ListenAndServe(":8080", nil)
}
