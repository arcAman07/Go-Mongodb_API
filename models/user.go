package models
type User struct {
	Name string `json:"name,omitempty" bson:"name"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
}