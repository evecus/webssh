package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	sshtunnel "github.com/yourusername/webssh/internal/ssh"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Type       string `json:"type"`
	Data       string `json:"data,omitempty"`
	Host       string `json:"host,omitempty"`
	Port       int    `json:"port,omitempty"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	PrivateKey string `json:"private_key,omitempty"`
	Passphrase string `json:"passphrase,omitempty"`
	Rows       uint32 `json:"rows,omitempty"`
	Cols       uint32 `json:"cols,omitempty"`
}

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/ws", wsHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(indexHTML))
}

type wsConn struct {
	mu sync.Mutex
	c  *websocket.Conn
}

func (w *wsConn) writeJSON(v interface{}) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.c.WriteJSON(v)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket upgrade error: %v", err)
		return
	}
	defer conn.Close()

	wsc := &wsConn{c: conn}

	_, raw, err := conn.ReadMessage()
	if err != nil {
		return
	}

	var msg Message
	if err := json.Unmarshal(raw, &msg); err != nil {
		wsc.writeJSON(map[string]string{"type": "error", "data": "invalid request"})
		return
	}

	if msg.Port == 0 {
		msg.Port = 22
	}

	cfg := sshtunnel.Config{
		Host:       msg.Host,
		Port:       msg.Port,
		Username:   msg.Username,
		Password:   msg.Password,
		PrivateKey: []byte(msg.PrivateKey),
		Passphrase: []byte(msg.Passphrase),
	}

	session, err := sshtunnel.Connect(cfg)
	if err != nil {
		wsc.writeJSON(map[string]string{"type": "error", "data": err.Error()})
		return
	}
	defer session.Close()

	wsc.writeJSON(map[string]string{"type": "connected"})

	outCh := make(chan []byte, 128)
	errCh := make(chan error, 2)
	session.ReadLoop(outCh, errCh)

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case data, ok := <-outCh:
				if !ok {
					return
				}
				if err := wsc.writeJSON(map[string]string{"type": "output", "data": string(data)}); err != nil {
					return
				}
			case sshErr := <-errCh:
				if sshErr != nil {
					wsc.writeJSON(map[string]string{"type": "error", "data": sshErr.Error()})
				} else {
					wsc.writeJSON(map[string]string{"type": "closed"})
				}
				return
			}
		}
	}()

	for {
		select {
		case <-done:
			return
		default:
		}

		_, raw, err := conn.ReadMessage()
		if err != nil {
			return
		}

		var wsMsg Message
		if err := json.Unmarshal(raw, &wsMsg); err != nil {
			continue
		}

		switch wsMsg.Type {
		case "input":
			if _, err := session.Write([]byte(wsMsg.Data)); err != nil {
				return
			}
		case "resize":
			rows := wsMsg.Rows
			cols := wsMsg.Cols
			if rows == 0 {
				rows = 24
			}
			if cols == 0 {
				cols = 80
			}
			session.Resize(rows, cols)
		}
	}
}
