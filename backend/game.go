package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type clientMessage struct {
	Type  string `json:"type"`
	Index *int   `json:"index,omitempty"`
}

type serverMessage struct {
	Type  string      `json:"type"`
	State interface{} `json:"state,omitempty"`
	Error string      `json:"error,omitempty"`
}

type wsConn struct {
	conn     *websocket.Conn
	playerID string
	writeMu  sync.Mutex
}

func (c *wsConn) writeJSON(v interface{}) error {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()
	return c.conn.WriteJSON(v)
}

type gameHub struct {
	mu    sync.Mutex
	conns map[*wsConn]struct{}
}

var (
	hubsMu sync.Mutex
	hubs   = map[string]*gameHub{}
)

func getOrCreateHub(gameID string) *gameHub {
	hubsMu.Lock()
	defer hubsMu.Unlock()
	if hub, ok := hubs[gameID]; ok {
		return hub
	}
	hub := &gameHub{conns: make(map[*wsConn]struct{})}
	hubs[gameID] = hub
	return hub
}

func getHub(gameID string) (*gameHub, bool) {
	hubsMu.Lock()
	defer hubsMu.Unlock()
	hub, ok := hubs[gameID]
	return hub, ok
}

func (h *gameHub) add(c *wsConn) {
	h.mu.Lock()
	h.conns[c] = struct{}{}
	h.mu.Unlock()
}

func (h *gameHub) remove(c *wsConn) {
	h.mu.Lock()
	delete(h.conns, c)
	h.mu.Unlock()
}

func (h *gameHub) snapshot() []*wsConn {
	h.mu.Lock()
	defer h.mu.Unlock()
	out := make([]*wsConn, 0, len(h.conns))
	for c := range h.conns {
		out = append(out, c)
	}
	return out
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// gameWSHandler -> GET /ws?game_id=...&player_id=...
func gameWSHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	gameID := r.URL.Query().Get("game_id")
	playerID := r.URL.Query().Get("player_id")
	if gameID == "" || playerID == "" {
		http.Error(w, "missing game_id or player_id", http.StatusBadRequest)
		return
	}

	gamesMu.RLock()
	game, ok := games[gameID]
	gamesMu.RUnlock()
	if !ok {
		http.Error(w, "game not found", http.StatusNotFound)
		return
	}

	game.Lock()
	if _, ok := game.Players[playerID]; !ok {
		game.Unlock()
		http.Error(w, "player not in game", http.StatusBadRequest)
		return
	}
	game.Unlock()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket upgrade failed: %v", err)
		return
	}
	ws := &wsConn{conn: conn, playerID: playerID}
	hub := getOrCreateHub(gameID)
	hub.add(ws)
	defer func() {
		hub.remove(ws)
		_ = conn.Close()
	}()

	sendStateToConn(game, ws)

	for {
		var msg clientMessage
		if err := conn.ReadJSON(&msg); err != nil {
			return
		}
		switch msg.Type {
		case "play_card":
			if msg.Index == nil {
				_ = ws.writeJSON(serverMessage{Type: "error", Error: "missing index"})
				continue
			}
			if err := applyPlayCard(gameID, playerID, *msg.Index); err != nil {
				_ = ws.writeJSON(serverMessage{Type: "error", Error: err.Error()})
				continue
			}
			broadcastGameState(gameID)
		default:
			_ = ws.writeJSON(serverMessage{Type: "error", Error: "unknown action"})
		}
	}
}

func applyPlayCard(gameID, playerID string, index int) error {
	gamesMu.RLock()
	game, ok := games[gameID]
	gamesMu.RUnlock()
	if !ok {
		return errGameNotFound
	}

	game.Lock()
	defer game.Unlock()

	player, ok := game.Players[playerID]
	if !ok {
		return errPlayerNotInGame
	}
	if index < 0 || index >= len(player.Hand) {
		return errInvalidCardIndex
	}

	card := player.Hand[index]
	player.Hand = append(player.Hand[:index], player.Hand[index+1:]...)
	game.Table = append(game.Table, card)
	game.UpdatedAt = time.Now()
	return nil
}

func sendStateToConn(game *GameState, ws *wsConn) {
	game.Lock()
	state := buildGameState(game, ws.playerID)
	game.Unlock()
	_ = ws.writeJSON(serverMessage{Type: "state", State: state})
}

func broadcastGameState(gameID string) {
	hub, ok := getHub(gameID)
	if !ok {
		return
	}
	conns := hub.snapshot()
	if len(conns) == 0 {
		return
	}

	gamesMu.RLock()
	game, ok := games[gameID]
	gamesMu.RUnlock()
	if !ok {
		return
	}

	game.Lock()
	states := make(map[*wsConn]interface{}, len(conns))
	for _, ws := range conns {
		states[ws] = buildGameState(game, ws.playerID)
	}
	game.Unlock()

	for ws, state := range states {
		_ = ws.writeJSON(serverMessage{Type: "state", State: state})
	}
}

func buildGameState(game *GameState, playerID string) map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = game.ID
	resp["players"] = func() map[string]interface{} {
		out := map[string]interface{}{}
		for pid, p := range game.Players {
			if playerID != "" && pid != playerID {
				out[pid] = map[string]interface{}{
					"id":         p.ID,
					"name":       p.Name,
					"hand_count": len(p.Hand),
				}
			} else {
				out[pid] = p
			}
		}
		return out
	}()
	resp["player_order"] = game.PlayerOrder
	resp["table"] = game.Table
	resp["updated_at"] = game.UpdatedAt
	return resp
}

var (
	errGameNotFound     = &wsError{"game not found"}
	errPlayerNotInGame  = &wsError{"player not in game"}
	errInvalidCardIndex = &wsError{"invalid card index"}
)

type wsError struct {
	msg string
}

func (e *wsError) Error() string {
	return e.msg
}
