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
		if u.Message != nil {
			h.Data.SaveMessage(u.Message)
		}

		if u.ID == 0 {
			continue
		}
		switch {
		case u.Message != nil:
			h.Logger.LogMessage() <- u.Message
		case u.NewChatCreated != nil:
			h.Logger.LogNewChat() <- u.NewChatCreated
		case u.Chat != nil:
			h.Data.SetChat(u.Chat)
		case u.User != nil:
			h.Data.SetUser(u.User)
		case u.Self != nil:
			h.Data.SetSelf(u.Self)
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
