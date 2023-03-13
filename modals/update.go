package modals

type Update struct {
	// ID is the update id unique to user
	ID uint64 `json:"id"`
	// Message is the new incoming message
	Message *Message `json:"message"`
	// EditedMessage new version of message that was already sent
	EditedMessage *Message `json:"edited_message"`
	// User is the user data requested
	User *User `json:"user"`
	// Self is the user data of the client
	Self *User `json:"self"`
	// Chat is the chat data requested
	Chat *Chat `json:"chat"`
	// Error is the error message
	Error string `json:"error"`
}