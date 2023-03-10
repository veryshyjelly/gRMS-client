package logger

import (
	"fmt"
	"gRMS-client/data"
	"gRMS-client/modals"
	"log"

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
	Data     data.DataHandler
}

func NewChatLogger(data data.DataHandler) ChatLogger {
	return &MyChatLogger{
		Messages: make(chan *modals.Message, 100),
		NewChat:  make(chan *modals.Chat, 20),
		Error:    make(chan string, 10),
		Data:     data,
	}
}

func (c *MyChatLogger) StartLogging() {
	blue := color.FgBlue
	yellow := color.FgYellow
	red := color.FgRed
	green := color.FgGreen

	for {
		select {
		case m := <-c.Messages:
			fmt.Print("\033[1000D\033[K")
			if m == nil {
				log.Fatalln("message is nil")
			}

			from := c.Data.GetUser(m.From)
			if from == nil {
				from = &modals.User{Username: "unknown"}
			}
			chat := c.Data.GetChat(m.Chat)
			if chat == nil {
				chat = &modals.Chat{Title: "unknown"}
			}

			yellow.Printf("[%d] ", m.ID)
			green.Light().Printf("%s", from.Username)
			blue.Light().Printf("@(%s)", chat.Title)
			green.Printf(">> ")

			// sbuilder.WriteString(fmt.Sprintf("[%d] %s @ %s >>", m.ID,
			// from.Username, chat.Title))

			switch {
			case m.Text != nil:
				fmt.Print(*m.Text)
			case m.Photo != 0:
				fmt.Printf("photo(id:%d)", m.Photo)
			case m.Video != 0:
				fmt.Printf("video(id:%d)", m.Video)
			case m.Document != 0:
				fmt.Printf("document(id:%d)", m.Document)
			case m.Audio != 0:
				fmt.Printf("audio(id:%d)", m.Audio)
			case m.Animation != 0:
				fmt.Printf("animation(id:%d)", m.Animation)
			}

		case c := <-c.NewChat:
			fmt.Print("\033[1000D\033[K")
			if c == nil {
				log.Fatalln("chat is nil")
			}
			blue.Light().Printf("New Chat - %s ", c.Title)
			fmt.Printf("(id:%d)", c.ID)

		case e := <-c.Error:
			fmt.Print("\033[1000D\033[K")
			red.Printf("%v", e)
		}
		fmt.Printf("\n%s> ", Prompt)
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
