package main

import (
	"math/rand"
	"net/http"
)

type NewGameHandler struct{}

func (h NewGameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := rand.Intn(8)
	GAMES[id] = GameState{}
}
