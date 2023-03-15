package logger

import (
	"fmt"
	"gRMS-client/data"
	"gRMS-client/modals"

	"github.com/gookit/color"
)

var Prompt string = "~"

type ChatLogger interface {
	StartLogging()
	LogMessage() chan *modals.Message
	LogNewChat() chan *modals.Chat
	LogError() chan string
}

type MyChatLogger struct {
	Messages chan *modals.Message
	NewChat  chan *modals.Chat
	Error    chan string
	Data     data.Handler
}

func NewChatLogger(data data.Handler) ChatLogger {
	return &MyChatLogger{
		Messages: make(chan *modals.Message, 100),
		NewChat:  make(chan *modals.Chat, 20),
		Error:    make(chan string, 10),
		Data:     data,
	}
}

func (c *MyChatLogger) StartLogging() {
	blue := color.FgBlue
	red := color.FgRed

	for {
		select {
		case m := <-c.Messages:
			fmt.Print("\033[1000D\033[K")
			m.Log(c.Data.GetChat(m.Chat), c.Data.GetUser(m.From))

		case c := <-c.NewChat:
			fmt.Print("\033[1000D\033[K")
			blue.Light().Printf("New Chat - %s ", c.Title)
			fmt.Printf("(id:%d)\n", c.ID)

		case e := <-c.Error:
			fmt.Print("\033[1000D\033[K")
			red.Printf("%v\n", e)
		}

		fmt.Printf("%s> ", Prompt)
	}
}

func (c *MyChatLogger) LogMessage() chan *modals.Message {
	return c.Messages
}

func (c *MyChatLogger) LogNewChat() chan *modals.Chat {
	return c.NewChat
}

func (c *MyChatLogger) LogError() chan string {
	return c.Error
}