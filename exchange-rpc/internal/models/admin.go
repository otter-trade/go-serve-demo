package models

type Admin struct {
	Name      string `bson:"name"`
	Email     string `bson:"email"`
	CreatedAt string `bson:"created_at"`
	UpdatedAt string `bson:"updated_at"`
}
