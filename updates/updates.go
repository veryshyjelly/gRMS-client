package updates

import (
	"gRMS-client/data"
	logger "gRMS-client/log"
	"gRMS-client/modals"
)

type UpdatesHandler struct {
	Updates chan modals.Update
	Logger  logger.ChatLogger
	Data    data.DataHandler
}

func (h *UpdatesHandler) Start() {
	for {
		u := <-h.Updates
		switch {
		case u.Message != nil:
			h.Logger.LogMessage() <- u.Message
		case u.NewChatCreated != nil:
			h.Logger.LogNewChat() <- u.NewChatCreated
		case u.Chat != nil:
			h.Data.SetChat(u.Chat)
		case u.User != nil:
			h.Data.SetUser(u.User)
		}
	}
}

func NewUpdatesHandler(updates chan modals.Update, l logger.ChatLogger, data data.DataHandler) *UpdatesHandler {
	return &UpdatesHandler{
		Updates: updates,
		Logger:  l,
		Data:    data,
	}
}
