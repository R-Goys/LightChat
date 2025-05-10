package handle

import (
	"net/http"

	connection "github.com/R-Goys/LightChat/conn"
	"github.com/gorilla/websocket"
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userID")
	TOID := r.FormValue("To")
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	connection.Conns.Add(userID, conn)
	defer conn.Close()
	conn.WriteMessage(websocket.TextMessage, []byte("Welcome to LightChat!"))
	if TOID != "" {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				break
			}
			if ok := connection.Conns.SoloChat(userID, TOID, message); !ok {
				conn.WriteMessage(websocket.BinaryMessage, []byte("Error: User not found."))
			}
		}
	} else {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				break
			}
			connection.Conns.Broadcast(userID, message)
		}
	}
	w.Write([]byte("Connection closed...."))
	connection.Conns.Delete(userID)
}
