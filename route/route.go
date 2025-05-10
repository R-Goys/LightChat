package route

import (
	"fmt"
	"net/http"

	"github.com/R-Goys/LightChat/handle"
)

func InitRoute() {
	http.HandleFunc("/chat", handle.ChatHandler)
	fmt.Println("http init")
	http.ListenAndServe(":9999", nil)
}
