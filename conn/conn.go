package connection

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ServerConn struct {
	conn sync.Map
}

type message struct {
	From string
	Msg  string
	To   string
}

var Conns = ServerConn{}

func (s *ServerConn) Add(userID string, conn *websocket.Conn) {
	s.conn.Store(userID, conn)
}

func (s *ServerConn) Delete(userID string) {
	s.conn.Delete(userID)
}

func (s *ServerConn) Broadcast(From string, msg []byte) []error {
	errs := make([]error, 0)
	Msg := message{
		From: From,
		Msg:  string(msg),
		To:   "",
	}
	s.conn.Range(func(key, value interface{}) bool {
		if key == From {
			return true
		}
		conn := value.(*websocket.Conn)
		Msg.To = key.(string)
		err := conn.WriteJSON(Msg)
		if err != nil {
			errs = append(errs, err)
		}
		return true
	})
	return errs
}

func (s *ServerConn) SoloChat(From string, To string, msg []byte) bool {
	Msg := message{
		From: From,
		Msg:  string(msg),
		To:   To,
	}
	conn, ok := s.conn.Load(To)
	if !ok {
		return false
	}
	conn.(*websocket.Conn).WriteJSON(Msg)
	return true
}
