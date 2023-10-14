package port

import "net/http"

type ChatHandler interface {
	InitConversation(w http.ResponseWriter, r *http.Request)
}
