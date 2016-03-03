package main

import "testing"

func TestInsertColor(t *testing.T) {
	game := GameState{}
	game.InsertColor("black", 0)
	if game.Board[0][0] != "black" {
		t.Errorf("Placed token wasn't black. It was: %s", game.Board[0][0])
	}
}

func TestInsertTooMany(t *testing.T) {
	game := GameState{}
	for i := 0; i < 6; i++ {
		if !game.InsertColor("black", 0) {
			t.Errorf("Token wasn't inserted %i time.", i)
		}
	}
	if game.InsertColor("black", 0) {
		t.Errorf("Too many tokens went into the playfield column.")
	}
}
