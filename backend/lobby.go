package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// createGameHandler -> POST /create
func createGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := randomID()
	game := &GameState{
		ID:          id,
		Players:     make(map[string]*Player),
		PlayerOrder: []string{},
		Table:       []Card{},
		DrawPile:    newDeck(),
		DiscardPile: []Card{},
		UpdatedAt:   time.Now(),
	}

	gamesMu.Lock()
	games[id] = game
	gamesMu.Unlock()

	writeJSON(w, http.StatusCreated, map[string]string{"game_id": id})
}

type joinReq struct {
	GameID string `json:"game_id"`
	Name   string `json:"name"`
}

// joinGameHandler -> POST /join
func joinGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body joinReq
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	gamesMu.RLock()
	game, ok := games[body.GameID]
	gamesMu.RUnlock()
	if !ok {
		http.Error(w, "game not found", http.StatusNotFound)
		return
	}

	playerID := randomID()
	player := &Player{ID: playerID, Name: body.Name}

	game.Lock()
	// deal 5 cards (or remaining)
	n := 5
	if len(game.DrawPile) < n {
		n = len(game.DrawPile)
	}
	player.Hand = append(player.Hand, game.DrawPile[:n]...)
	game.DrawPile = game.DrawPile[n:]

	game.Players[playerID] = player
	game.PlayerOrder = append(game.PlayerOrder, playerID)
	game.UpdatedAt = time.Now()
	game.Unlock()

	writeJSON(w, http.StatusOK, map[string]string{"player_id": playerID})
}
