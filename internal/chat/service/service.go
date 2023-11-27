package service

import (
	"encoding/json"
	"hostel-service/internal/chat/adapter/repo"
	"hostel-service/internal/chat/domain"
	wspkg "hostel-service/pkg/websocket"
	"log"

	"github.com/gorilla/websocket"
)

type ChatService interface {
	ProcessMessage(userId string, ws *websocket.Conn)
}

type chatService struct {
}

func NewChatService() *chatService {
	return &chatService{}
}

func (s *chatService) ProcessMessage(userId string, ws *websocket.Conn) {
	if err := s.fetchMessages(ws); err != nil {
		log.Printf("err: %v", err)
		return
	}
	wspkg.GetCtrl().Reader(userId, ws, s.handleMessage)
}

func (s *chatService) fetchMessages(ws *websocket.Conn) error {
	msgs, err := repo.FetchMessage()
	if err != nil {
		return err
	}
	data, err := json.Marshal(msgs)
	if err != nil {
		return err
	}
	if err := wspkg.GetCtrl().Writer(ws, data, websocket.TextMessage); err != nil {
		return err
	}
	return nil
}

func (s *chatService) handleMessage(msgType int, data []byte) error {
	var payload domain.Payload
	if err := json.Unmarshal(data, &payload); err != nil {
		log.Printf("err: %v", err)
		return err
	}

	// Store database used concurrency
	go func() {
		if err := repo.InsertMessage(payload); err != nil {
			log.Printf("err: %v", err)
		}
	}()
	//

	// Send message to peer
	ctrl := wspkg.GetCtrl()
	ws := ctrl.GetWebSocket(payload.ToId)
	if ws == nil {
		return nil
	}
	data, err := json.Marshal([]domain.Payload{payload})
	if err != nil {
		return err
	}
	if err := ctrl.Writer(ws, data, msgType); err != nil {
		return err
	}
	return nil
}
