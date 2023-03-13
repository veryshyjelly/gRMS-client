package client

import "gRMS-client/modals"

func (c *client) GetChat(chatId uint64) error {
	var req = modals.Req{
		GetChat: chatId,
	}
	return c.Conn.WriteJSON(req)
}

func (c *client) CreateChat(title string, participants []string) error {
	var req = modals.Req{
		NewChat: &modals.NewChatQuery{
			Title:        title,
			Participants: participants,
		},
	}
	return c.Conn.WriteJSON(req)
}

func (c *client) AddToChat(chatId uint64, usernames []string) error {
	var req = modals.Req{
		ChatJoin: &modals.UserQuery{
			ChatID: chatId,
			Users:  usernames,
		},
	}
	return c.Conn.WriteJSON(req)
}

func (c *client) RemoveFromChat(chatId uint64, usernames []string) error {
	var req = modals.Req{
		ChatKick: &modals.UserQuery{
			ChatID: chatId,
			Users:  usernames,
		},
	}
	return c.Conn.WriteJSON(req)
}

func (c *client) LeaveChat(chatId uint64) error {
	var req = modals.Req{
		LeaveChat: chatId,
	}
	return c.Conn.WriteJSON(req)
}