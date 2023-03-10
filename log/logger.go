package logger

import (
	"fmt"
	"gRMS-client/data"
	"gRMS-client/modals"
	"log"
	"strings"
)

var Prompt string = "~"

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
		Messages: make(chan *modals.Message, 100),
		NewChat:  make(chan *modals.Chat, 20),
		Data:     data,
	}
}

func (c *MyChatLogger) StartLogging() {
	for {
		select {
		case m := <-c.Messages:
			if m == nil {
				log.Fatalln("message is nil")
			}
			var sbuilder = strings.Builder{}

			from := c.Data.GetUser(m.From)
			if from == nil {
				from = &modals.User{Username: "unknown"}
			}
			chat := c.Data.GetChat(m.Chat)
			if chat == nil {
				chat = &modals.Chat{Title: "unknown"}
			}

			fmt.Print("\033[A\n")
			sbuilder.WriteString(fmt.Sprintf("[%d] %s @ %s >>", m.ID,
				from.Username, chat.Title))

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

			fmt.Print(sbuilder.String())
			fmt.Printf("\n%s> ", Prompt)

		case c := <-c.NewChat:
			if c == nil {
				log.Fatalln("chat is nil")
			}
			fmt.Printf("New Chat: [%d] %s\n", c.ID, c.Title)
		}
	}
}

func (c *MyChatLogger) LogMessage() chan *modals.Message {
	return c.Messages
}

func (c *MyChatLogger) LogNewChat() chan *modals.Chat {
	return c.NewChat
}
