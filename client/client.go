package client

import (
	"fmt"
	"gRMS-client/modals"
	"net/url"

	"github.com/gorilla/websocket"
)

type Client interface {
	Connect(host, path string) error
	SendMessage(chatId uint64, text string, replyId uint64) error
	SendMedia(chatId uint64, fileId uint64, fileType string, replyId uint64) error
	ForwardMessage(fromChatId uint64, messId uint64, toChatId uint64) error
	DeleteMessage(chatId uint64, messId uint64) error
	CreateChat(title string, participants []string) error
	AddToChat(chatId uint64, usernames []string) error
	RemoveFromChat(chatId uint64, usernames []string) error
	LeaveChat(chatId uint64) error // Leave chat should not work in dms
	GetChat(chatId uint64) error
	GetUser(userId uint64) error
	GetSelf() error
	GetUpdatesChan() chan modals.Update
	ChangePassword()
	ChangeUsername()
	Close() error
}

func NewClient(username, password string) Client {
	return &MyClient{
		Username: username,
		Password: password,
		Updates:  make(chan modals.Update, 100),
	}
}

type MyClient struct {
	Username string
	Password string
	Debug    bool
	Updates  chan modals.Update
	Conn     *websocket.Conn
}

func (c *MyClient) Connect(host, path string) (err error) {
	URL := url.URL{Scheme: "ws", Host: host, Path: "ws",
		RawQuery: fmt.Sprintf("username=%s&password=%s", c.Username, c.Password)}
	c.Conn, _, err = websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		return err
	}

	for {
		var up modals.Update
		err := c.Conn.ReadJSON(&up)
		if err != nil {
			// log.Println("an error occurred", err)
			return nil
		}
		c.Updates <- up
	}
}

func (c *MyClient) GetUpdatesChan() chan modals.Update {
	return c.Updates
}

func (c *MyClient) GetSelf() error {
	var req = modals.Req{
		GetSelf: 1,
	}
	return c.Conn.WriteJSON(req)
}

func (c *MyClient) Close() error {
	return c.Conn.Close()
}
