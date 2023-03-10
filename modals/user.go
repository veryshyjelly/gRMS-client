package modals

type User struct {
	// Unique ID of the User
	ID uint64 `json:"id" gorm:"primaryKey"`
	// FirstName of the user
	FirstName string `json:"first_name" validate:"required,min=3"`
	// LastName of the user
	LastName string `json:"last_name"`
	// Username is unique username of the user
	Username string `json:"username" validate:"required,min=3"`
	// Bio of the user
	Bio string `json:"bio"`
	// Chats is the list of chats the user is in
	Chats []Participant `json:"chats" gorm:"foreignKey:UserID"`
}
