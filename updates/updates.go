package updates

import (
	"gRMS-client/data"
	logger "gRMS-client/log"
	"gRMS-client/modals"
)

type Handler struct {
	Updates chan modals.Update
	Logger  logger.ChatLogger
	Data    data.Handler
}

func (h *Handler) Start() {
	for {
		u := <-h.Updates
		if u.Message != nil {
			h.Data.SaveMessage(u.Message)
		}
		if u.Error != "" {
			h.Logger.LogError() <- u.Error
		}
		if u.Chat != nil {
			h.Data.SetChat(u.Chat)
		}
		if u.User != nil {
			h.Data.SetUser(u.User)
		}
		if u.Self != nil {
			h.Data.SetSelf(u.Self)
		}
		if u.ID == 0 {
			continue
		}
		if mess := u.Message; mess != nil {
			h.Logger.LogMessage() <- mess
			if mess.NewChatCreated {
				h.Logger.LogNewChat() <- u.Chat
			} else if mess.NewChatMember != 0 {

			}
		}
	}
}

func NewUpdatesHandler(updates chan modals.Update, l logger.ChatLogger, data data.Handler) *Handler {
	return &Handler{
		Updates: updates,
		Logger:  l,
		Data:    data,
	}
}