package main

type GameState struct {
	Board [6][7]string
}

func (gs *GameState) InsertColor(color string, column int) bool {
	for i := 0; i < 6; i++ {
		if gs.Board[i][column] == "" {
			gs.Board[i][column] = color
			return true
		}
	}
	return false
}
