package modals

type Chat struct {
	// Unique ID of the Chat
	ID uint64 `json:"id" gorm:"primaryKey"`
	// Title of the Chat
	Title string `json:"title"`
	// Members is the list of usernames in the chat
	Members []Participant `json:"members" gorm:"foreignKey:ChatID"`
	// Admins is the list of admins in the chat
	Admins []Admin `json:"admins" gorm:"foreignKey:ChatID"`
	// DP is the display picture of the chat
	DP uint64 `json:"dp"`
	// Description is the chat description
	Description string `json:"description"`
	// InviteLink is the current active invite link
	InviteLink string `json:"invite_link"`
}

type Admin struct {
	// Unique ID of the ChatId
	ID uint64 `json:"-" gorm:"primaryKey"`
	// UserID is the ID of the user
	UserID uint64 `json:"user_id"`
	// ChatID is the ID of the chat
	ChatID uint64 `json:"chat_id"`
}

type Participant struct {
	// Unique ID of the ChatId
	ID uint64 `json:"-" gorm:"primaryKey"`
	// UserID is the ID of the user
	UserID uint64 `json:"user_id"`
	// ChatID is the ID of the chat
	ChatID uint64 `json:"chat_id"`
}
