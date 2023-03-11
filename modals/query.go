package modals

type Req struct {
	Message        *MessQuery    `json:"message"`
	NewChat        *NewChatQuery `json:"new_chat"`
	ChatJoin       *UserQuery    `json:"add_user"`
	ChatKick       *UserQuery    `json:"kick_user"`
	Promote        *UserQuery    `json:"promote"`
	Demote         *UserQuery    `json:"demote"`
	GetUser        uint64        `json:"get_user"`
	GetChat        uint64        `json:"get_chat"`
	GetSelf        uint64        `json:"get_self"`
	LeaveChat      uint64        `json:"leave_chat"`
	ForwardMessage *ForwardQuery `json:"forward"`
}

type MessQuery struct {
	ChatID           uint64 `json:"chat_id"`
	Text             string `json:"text,omitempty"`
	Document         uint64 `json:"doc,omitempty"`
	Photo            uint64 `json:"photo,omitempty"`
	Audio            uint64 `json:"audio,omitempty"`
	Video            uint64 `json:"video,omitempty"`
	Animation        uint64 `json:"animation,omitempty"`
	Thumb            uint64 `json:"thumb,omitempty"`
	Caption          string `json:"caption,omitempty"`
	ReplyToMessageID uint64 `json:"reply_to_message_id,omitempty"`
}

type NewChatQuery struct {
	Title        string   `json:"title"`
	Participants []string `json:"participants"`
}

type UserQuery struct {
	ChatID uint64   `json:"chat_id"`
	Users  []string `json:"users"`
}

type ForwardQuery struct {
	ToChatId   uint64 `json:"to_chat"`
	FromChatId uint64 `json:"from_chat"`
	MessageId  uint64 `json:"mess_id"`
}
