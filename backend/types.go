package main

// types and helpers (put once in the package)
import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Card struct {
	Suit  string `json:"suit"`
	Value string `json:"value"`
}

type Player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Hand []Card `json:"hand"`
}

type GameState struct {
	sync.Mutex
	ID          string             `json:"id"`
	Players     map[string]*Player `json:"players"`
	PlayerOrder []string           `json:"player_order"`
	Table       []Card             `json:"table"`
	DrawPile    []Card             `json:"-"`
	DiscardPile []Card             `json:"-"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

var (
	gamesMu sync.RWMutex
	games   = make(map[string]*GameState)
)

func newDeck() []Card {
	suits := []string{"♠", "♥", "♦", "♣"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	deck := make([]Card, 0, 52)
	for _, s := range suits {
		for _, v := range values {
			deck = append(deck, Card{Suit: s, Value: v})
		}
	}
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func randomID() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
