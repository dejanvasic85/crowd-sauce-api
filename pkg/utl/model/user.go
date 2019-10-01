package model

// User represents a user model
type User struct {
	Base
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
}
