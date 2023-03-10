package client

import (
	"fmt"
	"gRMS-client/modals"
)

func (c *MyClient) SendMessage(chatId uint64, text string, replyId uint64) error {
	return c.Conn.WriteJSON(modals.Req{
		Message: &modals.MessQuery{
			ChatID:           chatId,
			Text:             text,
			ReplyToMessageID: replyId,
		},
	})
}

func (c *MyClient) SendMedia(chatId uint64, fileId uint64, fileType string, replyId uint64) error {
	var req = modals.Req{
		Message: &modals.MessQuery{
			ChatID: chatId,
		},
	}

	switch fileType {
	case "photo":
		req.Message.Photo = fileId
	case "video":
		req.Message.Video = fileId
	case "audio":
		req.Message.Audio = fileId
	case "document":
		req.Message.Document = fileId
	case "animation":
		req.Message.Animation = fileId
	default:
		return fmt.Errorf("invalid file type")
	}

	return c.Conn.WriteJSON(req)
}

func (c *MyClient) ForwardMessage(fromChatId uint64, messId uint64, toChatId uint64) error {
	var req = modals.Req{
		ForwardMessage: &modals.ForwardQuery{
			ToChatId:   toChatId,
			FromChatId: fromChatId,
			MessageId:  messId,
		},
	}
	return c.Conn.WriteJSON(req)
}

func (c *MyClient) DeleteMessage(chatId uint64, messId uint64) error {
	return nil
}
