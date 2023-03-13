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
	LogNewMember() chan *modals.Message
	LogError() chan string
}

type chatLogger struct {
	Messages  chan *modals.Message
	NewMember chan *modals.Message
	NewChat   chan *modals.Chat
	Error     chan string
	Data      data.Handler
}

func NewChatLogger(data data.Handler) ChatLogger {
	return &chatLogger{
		Messages:  make(chan *modals.Message, 100),
		NewChat:   make(chan *modals.Chat, 20),
		NewMember: make(chan *modals.Message, 20),
		Error:     make(chan string, 10),
		Data:      data,
	}
}

func (c *chatLogger) StartLogging() {
	blue := color.FgBlue
	red := color.FgRed

	for {
		select {
		case m := <-c.Messages:
			fmt.Print("\033[1000D\033[K")
			m.Log(c.Data.GetChat(m.Chat), c.Data.GetUser(m.From))

		case c := <-c.NewChat:
			fmt.Print("\033[1000D\033[K")
			blue.Light().Printf("- New Chat : %s ", c.Title)
			fmt.Printf("(id:%d)\n", c.ID)

		case e := <-c.Error:
			fmt.Print("\033[1000D\033[K")
			red.Printf("%v\n", e)

		case m := <-c.NewMember:
			fmt.Print("\033[1000D\033[K")
			blue.Printf("- New Member : %s@(%s)", c.Data.GetUser(m.NewChatMember).Username, c.Data.GetChat(m.Chat).Title)
		}

		fmt.Printf("%s> ", Prompt)
	}
}

func (c *chatLogger) LogMessage() chan *modals.Message {
	return c.Messages
}

func (c *chatLogger) LogNewChat() chan *modals.Chat {
	return c.NewChat
}

func (c *chatLogger) LogNewMember() chan *modals.Message {
	return c.NewMember
}

func (c *chatLogger) LogError() chan string {
	return c.Error
}