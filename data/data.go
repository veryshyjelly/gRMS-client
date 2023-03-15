package data

import (
	"gRMS-client/client"
	"gRMS-client/modals"
	"log"
	"time"
)

type Handler interface {
	SetSelf(u *modals.User)
	GetSelf() *modals.User
	GetUser(uId uint64) *modals.User
	GetChat(uId uint64) *modals.Chat
	SetUser(user *modals.User)
	SetChat(chat *modals.Chat)
	SaveMessage(mess *modals.Message)
	GetMessages(chatID uint64) []*modals.Message
}

type handler struct {
	Self     *modals.User
	Users    map[uint64]*modals.User
	Chats    map[uint64]*modals.Chat
	Messages map[uint64][]*modals.Message
	Client   client.Client
}

func NewDataHandler(c client.Client) Handler {
	return &handler{
		Users:    make(map[uint64]*modals.User),
		Chats:    make(map[uint64]*modals.Chat),
		Messages: make(map[uint64][]*modals.Message),
		Client:   c,
	}
}

func (d *handler) GetUser(uId uint64) *modals.User {
	if u, ok := d.Users[uId]; ok {
		return u
	} else {
		err := d.Client.GetUser(uId)
		if err != nil {
			log.Println("error getting user:", err)
			return nil
		}
		time.Sleep(time.Millisecond * 200)
		if u, ok := d.Users[uId]; ok {
			return u
		} else {
			return nil
		}
	}
}

func (d *handler) SetUser(user *modals.User) {
	d.Users[user.ID] = user
}

func (d *handler) GetChat(cId uint64) *modals.Chat {
	if c, ok := d.Chats[cId]; ok {
		return c
	} else {
		err := d.Client.GetChat(cId)
		if err != nil {
			log.Println("error getting chat:", err)
			return nil
		}
		time.Sleep(time.Millisecond * 200)
		if c, ok := d.Chats[cId]; ok {
			return c
		} else {
			return nil
		}
	}
}

func (d *handler) SetChat(chat *modals.Chat) {
	d.Chats[chat.ID] = chat
	for _, v := range chat.Members {
		if _, ok := d.Users[v.UserID]; !ok {
			d.Client.GetUser(v.UserID)
		}
	}
}

func (d *handler) SaveMessage(mess *modals.Message) {
	if _, ok := d.Chats[mess.Chat]; !ok {
		d.Client.GetChat(mess.Chat)
	}

	if c, ok := d.Messages[mess.Chat]; ok {
		d.Messages[mess.Chat] = append(c, mess)
	} else {
		d.Messages[mess.Chat] = []*modals.Message{mess}
	}
}

func (d *handler) GetMessages(chatID uint64) []*modals.Message {
	if v, ok := d.Messages[chatID]; ok {
		return v
	}

	return []*modals.Message{}
}

func (d *handler) SetSelf(u *modals.User) {
	d.Self = u
}

func (d *handler) GetSelf() *modals.User {
	if d.Self == nil {
		d.Client.GetSelf()
	}
	time.Sleep(time.Millisecond * 200)
	return d.Self
}