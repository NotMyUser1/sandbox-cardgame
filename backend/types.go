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
	ID      string `json:"id"`
	ImageID string `json:"image_id"`
}

type Player struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Hand   []Card   `json:"hand"`
	Tricks [][]Card `json:"tricks"`
}

func (p *Player) toApi() *PlayerApi {
	return &PlayerApi{
		ID:         p.ID,
		Name:       p.Name,
		HandCount:  len(p.Hand),
		TrickCount: len(p.Tricks),
	}
}

type PlayerApi struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	HandCount  int    `json:"hand_count"`
	TrickCount int    `json:"trick_count"`
}

type GameState struct {
	sync.Mutex
	ID                   string
	Players              map[string]*Player
	PlayerOrder          []string
	Table                []Card
	DrawPile             []Card
	DiscardPile          []Card
	UpdatedAt            time.Time
	InitialHandCardCount int
	RevealCards          bool
}

func (g *GameState) toAPI(playerID string, reveal bool) GameStateApi {
	result := GameStateApi{
		ID:          g.ID,
		Player:      *g.Players[playerID],
		Players:     make(map[string]*any, len(g.Players)-1),
		PlayerOrder: g.PlayerOrder,
		Table:       g.Table,
		UpdatedAt:   g.UpdatedAt,
	}
	for key, player := range g.Players {
		if key != playerID {
			playerApi := any(*player.toApi())
			if reveal {
				playerApi = any(*player)
			}
			result.Players[key] = &playerApi
		}
	}
	return result
}

type GameStateApi struct {
	ID          string          `json:"id"`
	Player      Player          `json:"me"`
	Players     map[string]*any `json:"players"`
	PlayerOrder []string        `json:"player_order"`
	Table       []Card          `json:"table"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

var (
	gamesMu sync.RWMutex
	games   = make(map[string]*GameState)
)

func newDeck() []Card {
	suits := []string{"♠", "♥", "♦", "♣"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	deck := make([]Card, 0, 52)
	idSet := make(map[string]struct{}, 52)
	for _, s := range suits {
		for _, v := range values {
			id := randomID()
			for _, idAlreadyUsed := idSet[id]; idAlreadyUsed; {
				id = randomID()
				_, idAlreadyUsed = idSet[id]
			}
			deck = append(deck, Card{ID: randomID(), ImageID: s + v})
			idSet[id] = struct{}{}
		}
	}
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}

func writeJSON(w http.ResponseWriter, status int, v map[string]string) {
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
