package handler

import (
	"hostel-service/internal/chat/port"
	"hostel-service/internal/chat/service"
	wspkg "hostel-service/internal/package/websocket"
	"log"
	"net/http"
)

type chatHandler struct {
	chatService service.ChatService
}

func NewChatHandler(chatService service.ChatService) port.ChatHandler {
	return &chatHandler{
		chatService: chatService,
	}
}

func (h *chatHandler) InitConversation(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	if userId == "" {
		log.Println("require user id")
	}
	ws, err := wspkg.GetCtrl().CreateWS(userId, w, r)
	if err != nil {
		log.Printf("err: %v", err)
	}
	h.chatService.ProcessMessage(userId, ws)
}
