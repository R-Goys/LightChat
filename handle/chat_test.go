package handle

import (
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestChat(t *testing.T) {
	dialer := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}

	// 注意：URL 一定要带 ws://
	conn, _, err := dialer.Dial("ws://localhost:9999/chat", nil)
	if err != nil {
		t.Fatal("连接失败:", err)
	}
	defer conn.Close()

	// 向服务器发送一个简单的消息
	err = conn.WriteMessage(websocket.TextMessage, []byte("ping"))
	if err != nil {
		t.Fatal("向服务器发送消息失败", err)
		return
	}
	t.Log("连接成功")
}
