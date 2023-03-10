package data

import (
	"gRMS-client/client"
	"gRMS-client/modals"
	"log"
	"time"
)

type DataHandler interface {
	GetUser(uId uint64) *modals.User
	GetChat(uId uint64) *modals.Chat
	SetUser(user *modals.User)
	SetChat(chat *modals.Chat)
}

type MyDataHandler struct {
	Users  map[uint64]*modals.User
	Chats  map[uint64]*modals.Chat
	Client client.Client
}

func NewDataHandler(c client.Client) DataHandler {
	return &MyDataHandler{
		Users:  make(map[uint64]*modals.User),
		Chats:  make(map[uint64]*modals.Chat),
		Client: c,
	}
}

func (d *MyDataHandler) GetUser(uId uint64) *modals.User {
	if u, ok := d.Users[uId]; ok {
		return u
	} else {
		err := d.Client.GetUser(uId)
		if err != nil {
			log.Println("error getting user:", err)
			return nil
		}
		time.Sleep(time.Millisecond * 100)
		return d.Users[uId]
	}
}

func (d *MyDataHandler) SetUser(user *modals.User) {
	d.Users[user.ID] = user
}

func (d *MyDataHandler) GetChat(cId uint64) *modals.Chat {
	if c, ok := d.Chats[cId]; ok {
		return c
	} else {
		err := d.Client.GetChat(cId)
		if err != nil {
			log.Println("error getting chat:", err)
			return nil
		}
		time.Sleep(time.Millisecond * 100)
		return d.Chats[cId]
	}
}

func (d *MyDataHandler) SetChat(chat *modals.Chat) {
	d.Chats[chat.ID] = chat
}
