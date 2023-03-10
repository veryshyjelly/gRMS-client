package logger

import (
	"fmt"
	"gRMS-client/data"
	"gRMS-client/modals"
	"strings"
)

type ChatLogger interface {
	StartLogging()
	LogMessage() chan *modals.Message
	LogNewChat() chan *modals.Chat
}

type MyChatLogger struct {
	Messages chan *modals.Message
	NewChat  chan *modals.Chat
	Data     data.DataHandler
}

func NewChatLogger(data data.DataHandler) ChatLogger {
	return &MyChatLogger{
		Messages: make(chan *modals.Message),
		NewChat:  make(chan *modals.Chat),
		Data:     data,
	}
}

func (c *MyChatLogger) StartLogging() {
	select {
	case m := <-c.Messages:
		var sbuilder = strings.Builder{}

		sbuilder.WriteString(fmt.Sprintf("[%d] %s@%s >>", m.ID,
			c.Data.GetUser(m.From).Username, c.Data.GetChat(m.Chat).Title))

		switch {
		case m.Text != nil:
			sbuilder.WriteString(fmt.Sprintf(" %s ", *m.Text))
		case m.Photo != 0:
			sbuilder.WriteString(fmt.Sprintf(" photo(id:%d) ", m.Photo))
		case m.Video != 0:
			sbuilder.WriteString(fmt.Sprintf(" video(id:%d) ", m.Video))
		case m.Document != 0:
			sbuilder.WriteString(fmt.Sprintf(" document(id:%d) ", m.Document))
		case m.Audio != 0:
			sbuilder.WriteString(fmt.Sprintf(" audio(id:%d) ", m.Audio))
		case m.Animation != 0:
			sbuilder.WriteString(fmt.Sprintf(" animation(id:%d) ", m.Animation))
		}

		fmt.Println(sbuilder)

	case c := <-c.NewChat:
		fmt.Printf("New Chat: [%d] %s\n", c.ID, c.Title)
	}
}

func (c *MyChatLogger) LogMessage() chan *modals.Message {
	return c.Messages
}

func (c *MyChatLogger) LogNewChat() chan *modals.Chat {
	return c.NewChat
}
