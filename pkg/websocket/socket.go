package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketCtrl struct {
	upgrader websocket.Upgrader
	wsTable  map[string]*websocket.Conn
}

var wsCtrl *WebSocketCtrl

func GetCtrl() *WebSocketCtrl {
	if wsCtrl == nil {
		wsCtrl = &WebSocketCtrl{
			upgrader: websocket.Upgrader{
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
				CheckOrigin:     func(r *http.Request) bool { return true },
			},
			wsTable: map[string]*websocket.Conn{},
		}
	}
	return wsCtrl
}

func (ctrl *WebSocketCtrl) CreateWS(wsID string, w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {

	ws, err := ctrl.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	ctrl.wsTable[wsID] = ws
	return ws, nil
}

func (ctrl *WebSocketCtrl) GetTable() map[string]*websocket.Conn {
	return ctrl.wsTable
}

func (ctrl *WebSocketCtrl) GetWebSocket(key string) *websocket.Conn {
	if ws, ok := ctrl.wsTable[key]; ok {
		return ws
	}
	return nil
}

func (ctrl *WebSocketCtrl) Reader(userId string, conn *websocket.Conn, handlePayload func(int, []byte) error) {
	for {
		msgType, p, err := conn.ReadMessage()
		if _, ok := err.(*websocket.CloseError); ok {
			delete(ctrl.wsTable, userId)
			return
		}
		if err != nil {
			log.Printf("err: %v", err)
			return
		}
		if err := handlePayload(msgType, p); err != nil {
			log.Printf("err: %v", err)
		}
	}
}

func (ctrl *WebSocketCtrl) Writer(ws *websocket.Conn, data []byte, msgType int) error {
	if err := ws.WriteMessage(msgType, []byte(data)); err != nil {
		return err
	}
	return nil
}
